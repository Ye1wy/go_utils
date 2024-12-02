package pkg

import (
	"flag"
	"fmt"
	"go_utils/internal/find"
	"log"
)

func RunFind() {
	flag.Parse()

	path := "./"
	err := find.ValidingFlag()

	if err != nil {
		fmt.Printf("[Error] Flag error: %v\n", err)
		return
	}

	args := flag.Args()

	if len(args) > 0 {
		path = args[0]
	}

	entries, err := find.FilePathWalkDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e)
	}
}
