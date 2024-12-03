package pkg

import (
	"errors"
	"flag"
	"fmt"
	"go_utils/internal/wc"
	"strconv"
	"sync"
)

func RunWC(config *wc.Config) error {
	args := flag.Args()

	if len(args) < 1 {
		return errors.New("Nothing to read mate!")
	}

	var counterFunc func(string) (int, error)

	if config.ShowLine {
		counterFunc = wc.LineCounter

	} else if config.ShowWord {
		counterFunc = wc.WordsCounter

	} else if config.ShowChar {
		counterFunc = wc.CharCounter
	}

	run_count(args, counterFunc)

	return nil
}

func run_count(args []string, counterType func(string) (int, error)) {
	wg := sync.WaitGroup{}
	var mtx sync.Mutex
	var res []int

	for _, item := range args {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			temp, _ := counterType(fileName)

			mtx.Lock()
			res = append(res, temp)
			mtx.Unlock()

		}(item)

	}

	wg.Wait()

	for i, j := len(res)-1, 0; i >= 0; i-- {
		fmt.Println(strconv.Itoa(res[i]) + " " + args[j])
		j++
	}
}
