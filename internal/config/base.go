package config

import (
	"errors"
)

type FlagInterface interface {
	Name() string
	Register()
	ContainedValue() interface{}
}

type Config struct {
	Flags []FlagInterface
}

func (config *Config) AddFlag(flags FlagInterface) {
	config.Flags = append(config.Flags, flags)
	flags.Register()
}

func (config *Config) GetFlagValue(name string) (interface{}, error) {
	for _, item := range config.Flags {
		if item.Name() == name {
			return item.ContainedValue(), nil
		}
	}

	return nil, errors.New("that flag don't have in config")
}
