package wc

import (
	"errors"
	"flag"
)

var (
	lFlag = flag.Bool("l", false, "")
	mFlag = flag.Bool("m", false, "")
	wFlag = flag.Bool("w", false, "")
)

func ValidingFlag() (err error) {
	var flagCount int

	if *lFlag {
		flagCount++
	}

	if *mFlag {
		flagCount++
	}

	if *wFlag {
		flagCount++
	}

	if flagCount > 1 {
		err = errors.New("Only 1 flag")
		return
	}

	return nil
}
