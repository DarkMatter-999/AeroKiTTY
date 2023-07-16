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
		process: p
	}
	return &shell, nil

	/*
	p.Write([]byte("ls\r"))
	time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	_, err = p.Read(b)
	if err != nil {
		log.Printf("Failed to read pty %s", err)
	}
	log.Println(b)
	*/
}

