package aerokitty_core

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateTermWindow(filePath string) {
	app := app.New()
	window := app.NewWindow("AeroKiTTY | " + filePath)

	ui := widget.NewTextGrid()      
	ui.SetText("Welcome to AeroKiTTY")

	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(800, 600)),
			ui,
		),
	)

	spawnShell()


	window.ShowAndRun()
}

