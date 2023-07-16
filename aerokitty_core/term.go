package aerokitty_core

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/UserExistsError/conpty"
	"github.com/creack/pty"
)

func spawnShell() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" || runtime.GOOS == "freebsd" {
		c := exec.Command("sh")
		p, err := pty.Start(c)

		if err != nil {
			log.Printf("Failed to open pty %s", err)
			os.Exit(1)
		}

		defer c.Process.Kill()

		p.Write([]byte("ls\r"))
		time.Sleep(1 * time.Second)
		b := make([]byte, 1024)
		_, err = p.Read(b)
		if err != nil {
			log.Printf("Failed to read pty %s", err)
		}
		log.Println(b)

	} else if runtime.GOOS == "windows" {
		c, err := conpty.Start("cmd.exe")
		if err != nil {
			log.Printf("Failed to open conpty %s", err)
			os.Exit(1)
		}

		defer c.Close()

		c.Write([]byte("dir\r"))
		time.Sleep(1 * time.Second)
		b := make([]byte, 1024)
		_, err = c.Read(b)
		if err != nil {
			log.Printf("Failed to read pty %s", err)
		}
		log.Println(b)
	} else {
		log.Panic("Unsupported OS ", runtime.GOOS)
	}
}
