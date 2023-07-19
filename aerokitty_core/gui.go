package aerokitty_core

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateTermWindow(filePath string) {
	app := app.New()
	window := app.NewWindow("AeroKiTTY | " + filePath)

	label := widget.NewLabel("Welcome to AeroKiTTY")
    label.Wrapping = fyne.TextWrapWord

	shell, err := spawnShell()
	if err != nil {
		os.Exit(0)
	}

	onKeyType := func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
			_, _ = shell.Write([]byte{'\r'})
		} else if e.Name == fyne.KeyBackspace {
			_, _ = shell.Write([]byte{0x7f})
		}
	}

	onTypeChar := func(r rune) {
		_, _ = shell.Write([]byte{byte(r)})
	}

	window.Canvas().SetOnTypedKey(onKeyType)
	window.Canvas().SetOnTypedRune(onTypeChar)

	buffer := newTerm()
	reader := bufio.NewReader(shell)

	change := make(chan struct{})

	go func() {
        line := []rune{}
        buffer.buffer = append(buffer.buffer, line)
        previousBuffer := buffer.buffer

        for {
            r, _, err := reader.ReadRune()
            if err != nil {
                if err == io.EOF {
                    return
                }
                fyne.LogError("Error: %s", err)
                os.Exit(-1)
            }

            if r == 0x7F || r == 0x8 {
                if len(line) > 0 {
                    line = line[:len(line)-1]
                }
			} else if (r >= 32) && (r <= 126) || r == '\n' {
                line = append(line, r)
			} else {
				fmt.Printf("Got code %d", r)
			}

            buffer.buffer[len(buffer.buffer)-1] = line
            if r == '\n' {
                if len(buffer.buffer) > MaxBufferSize {
                    buffer.buffer = buffer.buffer[1:]
                }

                line = []rune{}
                buffer.buffer = append(buffer.buffer, line)
            }

            if !equalBuffers(previousBuffer, buffer.buffer) {
                previousBuffer = copyBuffer(buffer.buffer)
                change <- struct{}{} // Signal a buffer change
            }
        }
    }()

	/*
	go func() {
		for {
			//label.SetText("")
			var lines string
			for _, line := range buffer.buffer {
				lines = lines +  string(line)
			}
			label.SetText(string(lines))
			time.Sleep(10 * time.Millisecond)
		}
	}()
	*/

	go func() {
        for range change {
            lines := bufferToString(buffer.buffer)
			label.SetText(lines)
		}
    }()


	window.SetContent(
		fyne.NewContainerWithLayout(
            layout.NewGridWrapLayout(fyne.NewSize(800, 600)),
            label,
        ),	
	)

	window.ShowAndRun()
}

// Helper function to check if two buffers are equal
func equalBuffers(buf1, buf2 [][]rune) bool {
    if len(buf1) != len(buf2) {
        return false
    }

    for i := range buf1 {
        if !equalLines(buf1[i], buf2[i]) {
            return false
        }
    }

    return true
}

// Helper function to check if two lines are equal
func equalLines(line1, line2 []rune) bool {
    if len(line1) != len(line2) {
        return false
    }

    for i := range line1 {
        if line1[i] != line2[i] {
            return false
        }
    }

    return true
}

// Helper function to create a copy of the buffer
func copyBuffer(buffer [][]rune) [][]rune {
    copy := make([][]rune, len(buffer))
    for i, line := range buffer {
        copy[i] = make([]rune, len(line))
        copy[i] = append(copy[i], line...)
    }
    return copy
}

// Helper function to convert buffer to string
func bufferToString(buffer [][]rune) string {
    var lines string
    for _, line := range buffer {
        lines = lines + string(line)
    }
    return lines
}
