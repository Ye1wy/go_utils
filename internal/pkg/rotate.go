package pkg

import (
	"flag"
	"fmt"
	"go_utils/internal/config"
	"go_utils/internal/rotate"
	"sync"
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

	pathToComminDir := getFlagValue.(string)
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("[Error] Nothing to archivate, mate")
		return
	}

	var wg sync.WaitGroup

	for _, item := range args {
		wg.Add(1)

		go rotate.ProcessFile(item, pathToComminDir, &wg)
	}

	wg.Wait()
}
