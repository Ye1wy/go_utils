package main

import (
	"flag"
	"fmt"
	"go_utils/internal/find"
	"go_utils/internal/pkg"
)

func main() {
	config, err := find.ValidingFlag()

	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}

	path := "./"

	if len(flag.Args()) > 0 {
		path = flag.Arg(0)
	}

	pkg.RunFind(config, path)
}
