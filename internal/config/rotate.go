package config

import "flag"

const Dir = "directory"

type DirectoryFlag struct {
	Directory *string
}

func (d *DirectoryFlag) Register() {
	d.Directory = flag.String("d", "", "Directory were puted all archive")
}

func (d *DirectoryFlag) ContainedValue() interface{} {
	return *d.Directory
}

func (d *DirectoryFlag) Name() string {
	return Dir
}
