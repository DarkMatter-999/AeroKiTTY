//go:build windows
// +build windows

package aerokitty_core

import (
	"log"

	"github.com/UserExistsError/conpty"
)

type Shell struct {
	process *conpty.ConPty
}

func spawnShell() (*Shell, error){
	c, err := conpty.Start("cmd.exe")
	if err != nil {
		log.Printf("Failed to open conpty %s", err)
		return nil, err
	}

	shell := Shell{
		process: c,
	}
	return &shell, nil
}

func (s *Shell) Write(b []byte) (int, error) {
	return s.process.Write(b)
}

func (s *Shell) Read(b []byte) (int, error) {
	return s.process.Read(b)
}
