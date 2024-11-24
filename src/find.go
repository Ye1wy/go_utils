package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

var (
	slFlag  = flag.String("sl", "nothing", "")
	dFlag   = flag.String("d", "nothing", "")
	fFlag   = flag.String("f", "nothing", "")
	extFlag = flag.String("ext", "nothing", "")
)

func FilePathWalkDir(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	flag.Parse()

	if *fFlag == "nothing" && *extFlag != "nothing" {
		fmt.Println("ext flag comes only with f flag, read the description better. BE BETTER")
		return
	}

	entries, err := FilePathWalkDir("./")

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e)
	}
}
