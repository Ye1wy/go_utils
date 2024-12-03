package find

import (
	"errors"
	"flag"
)

type Config struct {
	ShowDirs  bool
	ShowFiles bool
	ShowLinks bool
	FileExt   string
}

func ValidingFlag() (*Config, error) {
	showDirs := flag.Bool("d", false, "Output only directories")
	showFiles := flag.Bool("f", false, "Output only files")
	showLinks := flag.Bool("sl", false, "Output only symlinks")
	fileExt := flag.String("ext", "", "Filter files by extension (works with -f)")

	flag.Parse()

	if *fileExt != "" && !*showFiles {
		return nil, errors.New("flag 'ext' can only be used with '-f'! Go ahead!")
	}

	return &Config{ShowDirs: *showDirs, ShowFiles: *showFiles, ShowLinks: *showLinks, FileExt: *fileExt}, nil
}
