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

	/*
	c.Write([]byte("dir\r"))
	time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	_, err = c.Read(b)
	if err != nil {
		log.Printf("Failed to read pty %s", err)
	}
	log.Println(b)
	*/
}
