package pkg

import (
	"flag"
	"fmt"
	"go_utils/internal/xargs"
)

func RunXargs() {
	config := &xargs.Config{}

	config.AddFlag(&xargs.ParallelFlag{})
	flag.Parse()

	arguments, err := xargs.ReadArgs()
	if err != nil {
		fmt.Printf("[Error] Error from reading args: %v\n", err)
		return
	}

	// fmt.Println(arguments)
	getedValue, err := config.GetFlagValue("parallel")
	if err != nil {
		fmt.Printf("[Error] Error in geting flag value: %v\n", err)
		return
	}

	command := flag.Args()[0]
	baseArgs := flag.Args()[1:]
	size := getedValue.(int)

	done := make(chan struct{}, size)
	for i := 0; i < len(*arguments); i++ {
		go xargs.Execute(command, baseArgs, *arguments, done)

		if len(done) == size {
			<-done
		}
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
}
