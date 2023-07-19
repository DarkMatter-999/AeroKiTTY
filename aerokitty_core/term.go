package aerokitty_core

import (
)

const MaxBufferSize = 16

type Terminal struct {
	buffer [][]rune
}

func newTerm() *Terminal {
	return &Terminal{
		buffer: [][]rune{},
	}
}
