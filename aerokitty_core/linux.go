// +build aix darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris !windows

package aerokitty_core

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"

)

type Shell struct {
	process *os.File
}

func spawnShell() (*Shell, error){
	c := exec.Command("sh")
	p, err := pty.Start(c)

	if err != nil {
		log.Printf("Failed to open pty %s", err)
		return nil, err
	}

	shell := Shell{
		process: p,
	}
	return &shell, nil
}

func (s *Shell) Write(b []byte) (int, error) {
	return s.process.Write(b)
}

func (s *Shell) Read(b []byte) (int, error) {
	return s.process.Read(b)
}
