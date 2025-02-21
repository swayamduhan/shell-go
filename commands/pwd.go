package commands

import (
	"fmt"
	"os"
)

func HandleWorkingDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd: encountered error: %v\n", err)
		return
	}

	fmt.Println(pwd)
}
