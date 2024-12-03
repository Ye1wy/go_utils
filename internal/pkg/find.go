package pkg

import (
	"flag"
	"fmt"
	"go_utils/internal/find"
	"log/slog"
)

func RunFind(config *find.Config, path string) {
	args := flag.Args()

	if len(args) > 0 {
		path = args[0]
	}

	entries, err := find.FilePathWalkDir(config, path)

	if err != nil {
		slog.Error("Error ", "err", err)
		return
	}

	for _, e := range entries {
		fmt.Println(e)
	}
}
