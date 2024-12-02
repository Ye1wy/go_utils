package pkg

import (
	"flag"
	"fmt"
	"log"

	"go_utils/internal/wc"
)

func RunWC() {
	flag.Parse()

	err := wc.ValidingFlag()

	if err != nil {
		fmt.Println("[Error] Flag error: ", err)
		return
	}

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("Nothing to read mate!")
	}
}
