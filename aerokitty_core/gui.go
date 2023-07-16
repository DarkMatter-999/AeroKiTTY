package aerokitty_core

import (
	"log"
	"os"

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

	shell, err := spawnShell()
	if err != nil {
		os.Exit(0)
	}


	shell.Write([]byte("dir\r"))
	//time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	_, err = shell.Read(b)
	if err != nil {
		log.Printf("Failed to read pty %s", err)
	}
	log.Println(b)


	window.ShowAndRun()
}

