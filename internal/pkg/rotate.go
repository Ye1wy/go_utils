package pkg

import (
	"flag"
	"fmt"
	"go_utils/internal/config"
)

func RunRotate() {
	conf := config.Config{}

	conf.AddFlag(&config.DirectoryFlag{})
	flag.Parse()

	getFlagValue, err := conf.GetFlagValue(config.Dir)
	if err != nil {
		fmt.Printf("[Error] Error from geting flag value: %v\n", err)
		return
	}

	value := getFlagValue.(string)

	fmt.Println(value)
}
