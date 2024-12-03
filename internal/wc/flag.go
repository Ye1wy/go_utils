package wc

import (
	"errors"
	"flag"
)

type Config struct {
	ShowLine bool
	ShowChar bool
	ShowWord bool
}

func ValidingFlag() (*Config, error) {
	showLine := flag.Bool("l", false, "")
	showChar := flag.Bool("m", false, "")
	showWord := flag.Bool("w", false, "")

	flag.Parse()

	var flagCount int

	if *showLine {
		flagCount++
	}

	if *showChar {
		flagCount++
	}

	if *showWord {
		flagCount++
	}

	if flagCount > 1 {
		return nil, errors.New("Only 1 flag")
	}

	if flagCount == 0 {
		*showWord = true
	}

	return &Config{ShowLine: *showLine, ShowChar: *showChar, ShowWord: *showWord}, nil
}
