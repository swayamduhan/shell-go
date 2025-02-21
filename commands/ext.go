package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/swayamduhan/shell-go/utils"
)


func RunExternalCmd(tokens []string){
	if len(tokens) == 0 {
		return
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// check for bash if on a windows machine
		if path, found := utils.HasBash(); found {
			// uses git bash if installed
			cmd = exec.Command(path, "-c", strings.Join(tokens, " "))
		} else {
			// no bash, use basic
			cmd = exec.Command(tokens[0], tokens[1:]...)
		}
	} else {
		// basic for unix
		cmd = exec.Command(tokens[0], tokens[1:]...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error executing command: %v\n", tokens[0], err)
	}
}

