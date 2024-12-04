package xargs

import (
	"fmt"
	"os"
	"os/exec"
)

func ReadArgs() {
	args := os.Args

	cmd := exec.Command(args[1], args[2])
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(stdout))

	// if err := cmd.Run(); err != nil {
	// 	return
	// }
}
