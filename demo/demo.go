package demo

import (
	"errors"
	"math"

	"github.com/davecgh/go-spew/spew"
)

const debug = false

func NextInteger(i int) (int, error) {
	if debug {
		spew.Dump(i)
	}

	if i == math.MaxInt64 {
		return 0, errors.New("i is too big")
	}

	return i + 1, nil
}
