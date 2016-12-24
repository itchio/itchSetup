package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/itchio/itchSetup/setup"
	"github.com/kardianos/osext"
	"github.com/lxn/walk"
	ui "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

func getUserDirectory(csidl win.CSIDL) (string, error) {
	localPathPtr := make([]uint16, 65536+2)
	var hwnd win.HWND
	success := win.SHGetSpecialFolderPath(hwnd, &localPathPtr[0], csidl, true)
	if !success {
		return "", errors.New("Could not get folder path")
	}
	return syscall.UTF16ToString(localPathPtr), nil
}

const MarkerName = ".itch-marker"

var localPath, roamingPath, desktopPath string

func SetupMain() {
	var err error

	localPath, err = getUserDirectory(win.CSIDL_LOCAL_APPDATA)
	if err != nil {
		showError(err.Error(), nil)
		os.Exit(1)
	}

	roamingPath, err = getUserDirectory(win.CSIDL_APPDATA)
	if err != nil {
		showError(err.Error(), nil)
		os.Exit(1)
	}

	desktopPath, err = getUserDirectory(win.CSIDL_DESKTOP)
	if err != nil {
		showError(err.Error(), nil)
		os.Exit(1)
	}

	log.Println("AppData local path: ", localPath)
	log.Println("AppData roam' path: ", roamingPath)
	log.Println("Desktop path:       ", desktopPath)

	log.Println("Should start", appName, ", looking for versions")

	execFolder, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	entries, err := ioutil.ReadDir(execFolder)
	if err != nil {
		log.Printf("")
	}

	foundMarker := false
	dirs := []string{}

	for _, entry := range entries {
		if !entry.IsDir() {
			if entry.Name() == MarkerName {
				foundMarker = true
			}
			continue
		}

		if !strings.HasPrefix(entry.Name(), "app-") {
			continue
		}

		dirs = append(dirs, entry.Name())
	}

	installDir := filepath.Join(localPath, appName)

	if foundMarker {
		log.Println("Found marker")

		if *uninstall == true {
			log.Println("Uninstalling app")

			log.Println("Removing shortcut...")
			err = os.Remove(shortcutPath())
			if err != nil {
				log.Println("While removing full shortcut", err.Error())
				log.Println("(Note: shortcut errors aren't critical)")
			}

			log.Println("Removing full directory...")
			err = os.RemoveAll(installDir)
			if err != nil {
				log.Println("While removing full directory", err.Error())
				log.Println("(Note: any itchSetup.exe-related errors are expected)")
			}

			log.Println("Removing uninstaller entry...")
			err = RemoveUninstallerRegistryKey(appName)
			if err != nil {
				log.Println("While removing uninstaller entry", err.Error())
				log.Println("(Note: these aren't critical)")
			}
			return
		}

		log.Println("Launching app")
		sort.Strings(dirs)
		log.Printf("Sorted app dirs:\n%s", strings.Join(dirs, "\n"))
		if len(dirs) > 0 {
			first := dirs[0]
			cmd := exec.Command(filepath.Join(first, "itch.exe"))
			err = cmd.Run()
			if err != nil {
				showError(fmt.Sprintf("Encountered a problem whilst launching itch: %s", err.Error()), nil)
			}

			log.Printf("App quit")
			os.Exit(0)
		}
		return
	}

	if *uninstall == true {
		log.Printf("Asked to uninstall but couldn't find marker, just quitting")
		os.Exit(0)
	}

	log.Println("Showing install GUI")
	showInstallGUI(installDir)
}

func showInstallGUI(installDirIn string) {
	var installer *setup.Installer

	var ni *walk.NotifyIcon
	var installDirLabel *walk.LineEdit
	var pb *walk.ProgressBar
	var progressLabel *walk.Label
	var mw *walk.MainWindow
	var imageView *walk.ImageView
	var progressComposite, optionsComposite *walk.Composite

	installDir := installDirIn
	versionInstallDir := ""

	var source setup.InstallSource
	sourceChan := make(chan setup.InstallSource, 1)

	kickoffInstall := func() {
		kickErr := func() error {
			err := os.MkdirAll(installDir, 0755)
			if err != nil {
				return err
			}

			execPath, err := osext.Executable()
			if err != nil {
				return err
			}

			// copy ourselves in install directory
			copyErr := func() error {
				installedExecPath := filepath.Join(installDir, "itchSetup.exe")
				execWriter, err := os.Create(installedExecPath)
				if err != nil {
					log.Println("Couldn't open target for writing, maybe already running from install location?")
					log.Println("Not copying and carrying on...")
					return nil
				}
				defer execWriter.Close()

				execReader, err := os.OpenFile(execPath, os.O_RDONLY, 0)
				if err != nil {
					return err
				}
				defer execReader.Close()

				_, err = io.Copy(execWriter, execReader)
				if err != nil {
					return err
				}
				return nil
			}()
			if copyErr != nil {
				return copyErr
			}

			source := <-sourceChan

			versionInstallDir = filepath.Join(installDir, fmt.Sprintf("app-%s", source.Version))
			installer.Install(versionInstallDir)

			return nil
		}()
		if kickErr != nil {
			showError(kickErr.Error(), mw)
		}
	}

	pickInstallLocation := func() {
		dlg := new(walk.FileDialog)

		dlg.Title = "Choose where the itch app should be installed"
		dlg.FilePath = installDir

		if ok, err := dlg.ShowBrowseFolder(mw); err != nil {
			log.Println(fmt.Sprintf("ShowBrowserFolder error: %s", err.Error()))
		} else if !ok {
			// nothing picked
		} else {
			installDir = dlg.FilePath
			installDirLabel.SetText(installDir)
		}
	}

	imageWidth := 622
	imageHeight := 301

	controlsHeight := 120
	windowHeight := imageHeight + 158 // found by trial & error

	windowSize := ui.Size{
		Width:  imageWidth,
		Height: windowHeight,
	}

	err := ui.MainWindow{
		Title:   "itch Setup",
		MinSize: windowSize,
		MaxSize: windowSize,
		Size:    windowSize,
		Layout: ui.VBox{
			MarginsZero: true,
			SpacingZero: true,
		},
		Children: []ui.Widget{
			ui.ImageView{
				AssignTo: &imageView,
				MinSize:  ui.Size{Width: imageWidth, Height: imageHeight},
				MaxSize:  ui.Size{Width: imageWidth, Height: imageHeight},
			},
			ui.Composite{
				MinSize: ui.Size{Height: controlsHeight},
				Layout: ui.VBox{
					Margins: ui.Margins{
						Left:  30,
						Right: 30,
					},
				},
				Children: []ui.Widget{
					ui.VSpacer{},
					ui.Label{
						Text: "Welcome to the itch installer! Grab a drink, pick an install location and proceed.",
					},
					ui.VSpacer{},
					ui.Composite{
						Layout: ui.HBox{
							MarginsZero: true,
						},
						Children: []ui.Widget{
							ui.LineEdit{
								AssignTo:    &installDirLabel,
								Text:        installDir,
								ReadOnly:    true,
								ToolTipText: "Click to change the install location",
								OnMouseUp: func(x, y int, button walk.MouseButton) {
									pickInstallLocation()
								},
							},
							ui.PushButton{
								MaxSize: ui.Size{Width: 1},
								Text:    "Install now",
								OnClicked: func() {
									progressComposite.SetVisible(true)
									optionsComposite.SetVisible(false)

									go kickoffInstall()
								},
							},
						},
					},
					ui.VSpacer{},
				},
				AssignTo: &optionsComposite,
			},
			ui.Composite{
				MinSize: ui.Size{Height: controlsHeight},
				Layout: ui.VBox{
					Margins: ui.Margins{
						Left:  30,
						Right: 30,
					},
				},
				Children: []ui.Widget{
					ui.VSpacer{},
					ui.ProgressBar{
						MinValue: 0,
						MaxValue: 1000,
						AssignTo: &pb,
					},
					ui.VSpacer{Size: 10},
					ui.Label{
						Text:     "Starting installation...",
						AssignTo: &progressLabel,
					},
					ui.VSpacer{},
				},
				Visible:  false,
				AssignTo: &progressComposite,
			},
		},
		AssignTo: &mw,
		OnSizeChanged: func() {
			if mw == nil {
				return
			}
			// this is kinda crap UX, but resizing the window is really ugly
			mw.SetSize(walk.Size{Width: windowSize.Width, Height: windowSize.Height})
		},
	}.Create()

	// remove maximize button
	style := win.GetWindowLong(mw.Handle(), win.GWL_STYLE)
	style &^= win.WS_MAXIMIZEBOX
	// style &^= win.WS_THICKFRAME
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, style)

	if err != nil {
		log.Fatal(err)
	}

	ni, err = walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}

	// see itchSetup.rc
	ic, err := walk.NewIconFromResourceId(101)
	if err != nil {
		log.Println("Could not load icon, oh well")
	} else {
		ni.SetIcon(ic)
		mw.SetIcon(ic)
	}

	err = ni.SetVisible(true)
	if err != nil {
		log.Printf("Could not make notifyicon visible: %s", err.Error())
	}

	setInstallerImage(imageView)

	installer = setup.NewInstaller(setup.InstallerSettings{
		OnError: func(message string) {
			go showError(message, mw)
		},
		OnProgressLabel: func(label string) {
			progressLabel.SetText(label)
		},
		OnProgress: func(progress float64) {
			pb.SetValue(int(progress * 1000.0))
		},
		OnSource: func(sourceIn setup.InstallSource) {
			sourceChan <- sourceIn
			source = sourceIn
		},
		OnFinish: func() {
			markerPath := filepath.Join(installDir, MarkerName)
			markerWriter, err := os.Create(markerPath)
			if err != nil {
				log.Println("While creating marker", err)
				showError(err.Error(), mw)
				os.Exit(1)
			}
			err = markerWriter.Close()
			if err != nil {
				log.Println("While closing marker", err)
				showError(err.Error(), mw)
				os.Exit(1)
			}

			// this creates $installDir/app.ico
			err = CreateUninstallRegistryEntry(installDir, "itch", source.Version)
			if err != nil {
				log.Printf("While creating registry entry: %s", err.Error())
			}

			err = CreateShortcut(ShortcutSettings{
				ShortcutFilePath: shortcutPath(),
				TargetPath:       filepath.Join(installDir, "itchSetup.exe"),
				Description:      "The best way to play your itch.io games",
				IconLocation:     filepath.Join(installDir, "app.ico"),
				WorkingDirectory: filepath.Join(installDir),
			})

			if err != nil {
				log.Println("While creating shortcut", err)
				showError(err.Error(), mw)
				os.Exit(1)
			}

			ni.ShowInfo("itch", "The installation went well, itch is now starting up!")

			itchPath := filepath.Join(versionInstallDir, "itch.exe")
			cmd := exec.Command(itchPath)
			err = cmd.Start()
			if err != nil {
				showError(err.Error(), mw)
			}

			time.Sleep(2 * time.Second)
			os.Exit(0)
		},
	})

	centerWindow(mw.AsFormBase())

	mw.Run()
}

func showError(errMsg string, parent walk.Form) {
	var dlg *walk.Dialog

	res, err := ui.Dialog{
		Title:    "Something went wrong",
		MinSize:  ui.Size{Width: 350},
		Layout:   ui.VBox{},
		AssignTo: &dlg,
		Children: []ui.Widget{
			ui.Composite{
				Layout: ui.HBox{
					MarginsZero: true,
				},
				Children: []ui.Widget{
					ui.Label{
						Text: errMsg,
					},
					ui.HSpacer{},
				},
			},
			ui.VSpacer{Size: 10},
			ui.Composite{
				Layout: ui.HBox{
					MarginsZero: true,
				},
				Children: []ui.Widget{
					ui.HSpacer{},
					ui.PushButton{
						Text:    "Okay :(",
						MaxSize: ui.Size{Width: 1},
						OnClicked: func() {
							dlg.Close(0)
						},
					},
					ui.HSpacer{},
				},
			},
		},
	}.Run(parent)

	centerWindow(dlg.AsFormBase())

	if err != nil {
		log.Printf("Error in dialog: %s\n", err.Error())
	}
	log.Printf("Dialog res: %#v\n", res)

	os.Exit(0)
}

func shortcutPath() string {
	return filepath.Join(desktopPath, "itch.lnk")
}
