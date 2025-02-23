package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/swayamduhan/shell-go/utils"
)

func RunExternalCmd(inputCmd string){
	if inputCmd == "" {
		return
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// check for bash if on a windows machine
		if path, found := utils.HasBash(); found {
			// uses git bash if installed
			cmd = exec.Command(path, "-c", inputCmd)
		} else {
			// no bash, use basic
			cmd = exec.Command("cmd", "/C", inputCmd)
		}
	} else {
		// basic for unix
		cmd = exec.Command("sh", "-c", inputCmd)
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

