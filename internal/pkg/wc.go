package pkg

import (
	"errors"
	"flag"
	"fmt"
	"go_utils/internal/wc"
	"sync"
)

func RunWC(config *wc.Config) error {
	args := flag.Args()

	if len(args) < 1 {
		return errors.New("Nothing to read mate!")
	}

	wg := sync.WaitGroup{}
	var mtx sync.Mutex

	if config.ShowLine {
		var res []int

		for _, item := range args {
			wg.Add(1)
			go func(fileName string) {
				defer wg.Done()
				temp, _ := wc.LineCounter(fileName)

				mtx.Lock()
				res = append(res, temp)
				mtx.Unlock()

			}(item)

		}

		wg.Wait()

		for i := len(res) - 1; i >= 0; i-- {
			fmt.Println(res[i])
		}

	}

	return nil
}
