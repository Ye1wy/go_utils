package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
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

func main() {
	flag.Parse()

	err := ValidingFlag()

	if err != nil {
		fmt.Println("[Error] Flag error: ", err)
		return
	}

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("Nothing to read mate!")
	}
}
