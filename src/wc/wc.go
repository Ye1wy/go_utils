package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	lFlag = flag.Bool("l", false, "")
	mFlag = flag.Bool("m", false, "")
	wFlag = flag.Bool("w", false, "")
)

func FileProcessor(file string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(file)
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("Nothing to read mate!")
	}

	go FileProcessor("Papa")
	go FileProcessor("Mama")

	FileProcessor("Hello")
}
