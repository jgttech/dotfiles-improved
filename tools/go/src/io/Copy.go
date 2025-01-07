package io

import (
	"errors"
	"io"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/os"
	_os "os"
)

func Copy(from, to string) (int64, error) {
	fromExists := os.Exists(from)
	toExists := os.Exists(to)

	if !fromExists {
		return 0, errors.New("File not found: " + from)
	}

	if !toExists {
		return 0, errors.New("File not found: " + to)
	}

	fromStat := assert.Must(_os.Stat(from))
	fromIo := assert.Must(_os.Open(from))
	toIo := assert.Must(_os.Open(to))

	defer fromIo.Close()
	defer toIo.Close()

	written, err := io.CopyN(toIo, fromIo, fromStat.Size())

	if err != nil {
		return 0, err
	}

	return written, nil
}
