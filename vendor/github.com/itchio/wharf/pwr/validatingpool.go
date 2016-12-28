package pwr

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-errors/errors"
	"github.com/itchio/wharf/pwr/drip"
	"github.com/itchio/wharf/tlc"
	"github.com/itchio/wharf/wsync"
)

type OnFileValidatedFunc func(fileIndex int64)

// A ValidatingPool will check files against their hashes, but doesn't
// check directories or symlinks
type ValidatingPool struct {
	// required

	Pool wsync.WritablePool
	// Container must match Pool - may have different file indices than Signature.Container
	Container *tlc.Container
	Signature *SignatureInfo

	Wounds chan *Wound

	OnFileValidated OnFileValidatedFunc

	// private //

	hashGroups map[int64][]wsync.BlockHash
	sctx       *wsync.Context
}

var _ wsync.WritablePool = (*ValidatingPool)(nil)

// GetSize is a pass-through to the underlying Pool
func (vp *ValidatingPool) GetSize(fileIndex int64) int64 {
	return vp.Pool.GetSize(fileIndex)
}

// GetReader is a pass-through to the underlying Pool, it doesn't validate
func (vp *ValidatingPool) GetReader(fileIndex int64) (io.Reader, error) {
	return vp.GetReadSeeker(fileIndex)
}

// GetReadSeeker is a pass-through to the underlying Pool, it doesn't validate
func (vp *ValidatingPool) GetReadSeeker(fileIndex int64) (io.ReadSeeker, error) {
	return vp.Pool.GetReadSeeker(fileIndex)
}

// GetWriter returns a writer that checks hashes before writing to the underlying
// pool's writer. It tries really hard to be transparent, but does buffer some data,
// which means some writing is only done when the returned writer is closed.
func (vp *ValidatingPool) GetWriter(fileIndex int64) (io.WriteCloser, error) {
	if vp.hashGroups == nil {
		err := vp.makeHashGroups()
		if err != nil {
			return nil, errors.Wrap(err, 1)
		}
		vp.sctx = wsync.NewContext(int(BlockSize))
	}

	w, err := vp.Pool.GetWriter(fileIndex)
	if err != nil {
		return nil, errors.Wrap(err, 1)
	}

	hashGroup := vp.hashGroups[fileIndex]
	blockIndex := int64(0)
	file := vp.Container.Files[fileIndex]
	fileSize := file.Size

	validate := func(data []byte) error {
		weakHash, strongHash := vp.sctx.HashBlock(data)
		start := blockIndex * BlockSize
		size := ComputeBlockSize(fileSize, blockIndex)

		if blockIndex >= int64(len(hashGroup)) {
			if vp.Wounds == nil {
				err := fmt.Errorf("%s: too large (%d blocks, tried to look up hash %d)",
					file.Path, len(hashGroup), blockIndex)
				return errors.Wrap(err, 1)
			}

			vp.Wounds <- &Wound{
				Kind:  WoundKind_FILE,
				Index: fileIndex,
				Start: start,
				End:   start + size,
			}
		} else {
			bh := hashGroup[blockIndex]

			if bh.WeakHash != weakHash {
				if vp.Wounds == nil {
					err := fmt.Errorf("%s: at block %d, expected weak hash %x, got %x", file.Path, blockIndex, bh.WeakHash, weakHash)
					return errors.Wrap(err, 1)
				}

				vp.Wounds <- &Wound{
					Kind:  WoundKind_FILE,
					Index: fileIndex,
					Start: start,
					End:   start + size,
				}
			} else if !bytes.Equal(bh.StrongHash, strongHash) {
				if vp.Wounds == nil {
					err := fmt.Errorf("%s: at block %d, expected strong hash %x, got %x", file.Path, blockIndex, bh.StrongHash, strongHash)
					return errors.Wrap(err, 1)
				}

				vp.Wounds <- &Wound{
					Kind:  WoundKind_FILE,
					Index: fileIndex,
					Start: start,
					End:   start + size,
				}
			}

			if vp.Wounds != nil {
				vp.Wounds <- &Wound{
					Kind:  WoundKind_CLOSED_FILE,
					Index: fileIndex,
					Start: start,
					End:   start + size,
				}
			}
		}

		blockIndex++
		return nil
	}

	dw := &drip.Writer{
		Writer:   w,
		Buffer:   make([]byte, BlockSize),
		Validate: validate,
	}

	return dw, nil
}

func (vp *ValidatingPool) makeHashGroups() error {
	// see blockpool's validator for a slightly different take on this
	pathToFileIndex := make(map[string]int64)
	for fileIndex, f := range vp.Container.Files {
		pathToFileIndex[f.Path] = int64(fileIndex)
	}

	vp.hashGroups = make(map[int64][]wsync.BlockHash)
	hashIndex := int64(0)

	for _, f := range vp.Signature.Container.Files {
		fileIndex := pathToFileIndex[f.Path]

		if f.Size == 0 {
			// empty files have a 0-length shortblock for historical reasons.
			hashIndex++
			continue
		}

		numBlocks := ComputeNumBlocks(f.Size)
		vp.hashGroups[fileIndex] = vp.Signature.Hashes[hashIndex : hashIndex+numBlocks]
		hashIndex += numBlocks
	}

	if hashIndex != int64(len(vp.Signature.Hashes)) {
		err := fmt.Errorf("expected to have %d hashes in signature, had %d", hashIndex, len(vp.Signature.Hashes))
		return errors.Wrap(err, 1)
	}

	return nil
}

// Close closes the underlying pool (and its reader, if any)
func (vp *ValidatingPool) Close() error {
	return vp.Pool.Close()
}
