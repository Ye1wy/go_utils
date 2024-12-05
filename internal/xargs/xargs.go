package xargs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ReadArgs() (*[]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			input = append(input, strings.Fields(line)...)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &input, nil
}

func Execute(command string, baseArgs []string, args []string, done chan struct{}) {
	cmd := exec.Command(command, append(baseArgs, args...)...)
	cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("[Error] Error in runing cmd: %v...\n", err)
		return
	}

	done <- struct{}{}
}
