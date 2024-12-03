package pkg

import (
	"errors"
	"flag"
	"go_utils/internal/wc"
)

func RunWC(config *wc.Config) error {
	args := flag.Args()

	if len(args) < 1 {
		return errors.New("Nothing to read mate!")
	}

	if config.ShowLine {

	}
}
