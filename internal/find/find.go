package find

import (
	"io/fs"
	"os"
	"path/filepath"
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

func FlagProcessingF(config *Config, root string, path string, info fs.FileInfo) (files []string) {
	if !info.IsDir() {
		if config.FileExt != "" {
			if filepath.Ext(path) == "."+config.FileExt {
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

func FilePathWalkDir(config *Config, root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if config.ShowDirs {
			files = append(files, FlagProcessingD(root, path, info)...)
		}

		if config.ShowFiles {
			files = append(files, FlagProcessingF(config, root, path, info)...)
		}

		if config.ShowLinks {
			files = append(files, FlagProcessingSL(root, path)...)
		}

		if !config.ShowDirs && !config.ShowFiles && !config.ShowLinks {
			files = append(files, FlagProcessingD(root, path, info)...)
			files = append(files, FlagProcessingF(config, root, path, info)...)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
