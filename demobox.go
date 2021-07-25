package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	demoboxApp := app.New()
	mainWindow := demoboxApp.NewWindow("Demobox v0.1")

	toolbar := widget.NewToolbar(
		// Open Demobox bundle
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			var (
				selectedfiles fyne.URIReadCloser
				fileerror     error
			)
			dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
				selectedfiles = file
				err = fileerror
			}, mainWindow)
		}),

		widget.NewToolbarSeparator(),

		// Start VM
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() { startVM() }),

		// Pause VM
		widget.NewToolbarAction(theme.MediaStopIcon(), func() { pauseVM() }),

		// Reset VM
		widget.NewToolbarAction(theme.ErrorIcon(), func() { resetVM() }),

		widget.NewToolbarSeparator(),

		// Save screenshot to PNG
		widget.NewToolbarAction(theme.ViewFullScreenIcon(), func() { saveImage() }),

		// Record video to MP4
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() { saveVideo() }),

		widget.NewToolbarSeparator(),

		// View RAM and Registers
		widget.NewToolbarAction(theme.ComputerIcon(), func() {
			settingsWindow := demoboxApp.NewWindow("Machine state")
			settingsWindow.Resize(fyne.NewSize(600, 700))
			settingsWindow.CenterOnScreen()
			settingsWindow.Show()
		}),

		widget.NewToolbarSeparator(),

		// Settings
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			settingsWindow := demoboxApp.NewWindow("Settings")
			settingsWindow.Resize(fyne.NewSize(400, 700))
			settingsWindow.CenterOnScreen()
			settingsWindow.Show()
		}),
	)

	// Main
	content := container.NewBorder(toolbar, nil, nil, nil)
	mainWindow.SetContent(content)
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}

func startVM() {
	log.Println("Start VM")
}

func pauseVM() {
	log.Println("Pause VM")
}

func resetVM() {
	log.Println("Restart VM")
}

func saveImage() {
	log.Println("Save image")
}

func saveVideo() {
	log.Println("Save video")
}
