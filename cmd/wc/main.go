package main

import (
	"fmt"
	"go_utils/internal/pkg"
	"go_utils/internal/wc"
)

func main() {
	config, err := wc.ValidingFlag()

	if err != nil {
		fmt.Println("[Error] Flag error: ", err)
		return
	}

	pkg.RunWC(config)
}
