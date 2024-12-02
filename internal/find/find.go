package find

import (
	"errors"
	"flag"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	slFlag      = flag.Bool("sl", false, "Output only symlinks in path")
	dFlag       = flag.Bool("d", false, "Output only dir in path")
	fFlag       = flag.Bool("f", false, "Output only files in path")
	extFlag     = flag.String("ext", "nothing", "")
	allNotExist bool
)

func FlagProcessingD(root string, path string, info fs.FileInfo) (files []string) {
	if info.IsDir() {
		if path == root {
			return
		}

		files = append(files, root+path)
	}

	return
}

func FlagProcessingF(root string, path string, info fs.FileInfo) (files []string) {
	if !info.IsDir() {
		if path == root {
			return
		}

		if _, err := os.Readlink(path); err == nil {
			return
		}

		if *extFlag != "nothing" {
			fileExtention := filepath.Ext(path)
			NeededExtention := "." + *extFlag

			if fileExtention == NeededExtention {
				files = append(files, root+path)
			}

		} else {
			files = append(files, root+path)
		}
	}

	return
}

func FlagProcessingSL(root string, path string) (files []string) {
	info, err := os.Readlink(path)

	if err != nil {
		return
	}

	if _, err := os.Stat(path); err != nil {
		files = append(files, root+path+"->[broken]")
		return
	}

	files = append(files, root+path+"->"+info)

	return
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string

	if !(*fFlag) && !(*slFlag) {
		files = append(files, ".")
	}

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if *dFlag {
			files = append(files, FlagProcessingD(root, path, info)...)
		}

		if *fFlag {
			files = append(files, FlagProcessingF(root, path, info)...)
		}

		if *slFlag {
			files = append(files, FlagProcessingSL(root, path)...)
		}

		if allNotExist {
			files = append(files, FlagProcessingD(root, path, info)...)
			files = append(files, FlagProcessingF(root, path, info)...)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

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
