package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

var (
	slFlag      = flag.Bool("sl", false, "Output only symlinks in path")
	dFlag       = flag.Bool("d", false, "Output only dir in path")
	fFlag       = flag.Bool("f", false, "Output only files in path")
	extFlag     = flag.String("ext", "nothing", "")
	allNotExist bool
)

func FlagProcessingD(path string, info fs.FileInfo) (files []string) {
	if info.IsDir() {
		files = append(files, path)
	}

	return
}

func FlagProcessingF(path string, info fs.FileInfo) (files []string) {
	if !info.IsDir() {
		files = append(files, path)
	}

	return
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if *dFlag {
			files = append(files, FlagProcessingD(path, info)...)

		}

		if *fFlag {
			files = append(files, FlagProcessingF(path, info)...)
		}

		if allNotExist {
			files = append(files, FlagProcessingD(path, info)...)
			files = append(files, FlagProcessingF(path, info)...)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("=============out=================")
	fmt.Println(files)

	return files, nil
}

func ValidingFlag() error {
	var err error

	if !(*fFlag) && *extFlag != "nothing" {
		err = errors.New("error parse flag, flag ext only work with f flag")
		return err
	}

	if !(*fFlag) && !(*dFlag) && !(*slFlag) {
		allNotExist = true
	}

	return nil
}

func main() {
	flag.Parse()

	path := "./"

	err := ValidingFlag()

	if err != nil {
		fmt.Printf("[Error] Flag error: %v\n", err)
		return
	}

	args := flag.Args()

	if len(args) > 0 {
		path = args[0]
	}

	fmt.Println(path)

	entries, err := FilePathWalkDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e)
	}
}
