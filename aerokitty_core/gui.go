package aerokitty_core

import (
	"bufio"
	"io"
	"os"
	"time"

	//"time"

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

	shell, err := spawnShell()
	if err != nil {
		os.Exit(0)
	}

	onKeyType := func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
			_, _ = shell.Write([]byte{'\r'})
		}
	}

	onTypeChar := func(r rune) {
		_, _ = shell.Write([]byte{byte(r)})
	}

	window.Canvas().SetOnTypedKey(onKeyType)
	window.Canvas().SetOnTypedRune(onTypeChar)

	buffer := newTerm()
	reader := bufio.NewReader(shell)

	go func() {
		line := []rune{}
		buffer.buffer = append(buffer.buffer, line)

		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return
				}
				fyne.LogError("Error: %s", err)
				os.Exit(-1)
			}

			line = append(line, r)
			buffer.buffer[len(buffer.buffer) - 1] = line
			if r == '\n' {
				if len(buffer.buffer) > MaxBufferSize {
					buffer.buffer = buffer.buffer[1:]
				}

				line = []rune{}
				buffer.buffer = append(buffer.buffer, line)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			var lines string
			for _, line := range buffer.buffer {
				lines = lines +  string(line)
			}
			ui.SetText("")
			ui.SetText(string(lines))
		}
	}()


	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(800, 600)),
			ui,
		),
	)

	window.ShowAndRun()
}

