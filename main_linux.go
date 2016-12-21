package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fasterthanlime/itchSetup/setup"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	imageWidth := 622
	imageHeight := 301

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("itch Setup")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 18)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	win.Add(box)

	i, err := gtk.ImageNewFromFile("data/installer.png")
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	box.Add(i)

	pb, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}
	pb.SetMarginStart(30)
	pb.SetMarginEnd(30)
	box.Add(pb)

	l, err := gtk.LabelNew("Warming up...")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	box.Add(l)

	// Set the default window size.
	win.SetDefaultSize(imageWidth, imageHeight+260)

	win.SetPosition(gtk.WIN_POS_CENTER)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	installDir := filepath.Join(os.Getenv("HOME"), ".itch-experimental/prefix")

	installer := setup.NewInstaller(setup.InstallerSettings{
		OnProgress: func(progress float64) {
			glib.IdleAdd(func() {
				pb.SetFraction(progress)
			})
		},
		OnProgressLabel: func(label string) {
			glib.IdleAdd(func() {
				l.SetText(label)
			})
		},
		OnError: func(message string) {
			glib.IdleAdd(func() {
				l.SetText(message)
			})
		},
		OnFinish: func() {
			itchPath := filepath.Join(installDir, "itch")
			cmd := exec.Command(itchPath)
			err := cmd.Start()
			if err != nil {
				glib.IdleAdd(func() {
					l.SetText(err.Error())
				})
			}

			time.Sleep(2 * time.Second)
			gtk.MainQuit()
		},
	})

	installer.Install(installDir)

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
