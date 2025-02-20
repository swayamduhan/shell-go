package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/swayamduhan/shell-go/utils"
)

var cmds = map[string]bool{
	"echo" : true,
	"exit" : true,
	"type" : true,
	"pwd" : true,
}

func HandleEcho(tokens []string) {
	fmt.Println(strings.Join(tokens[1:], " "))
}

func HandleExit(tokens []string){
	if len(tokens) > 2 {
		fmt.Println("too many arguments")
		return
	}
	fmt.Fprint(os.Stdin, "exitted shell\n")
	os.Exit(0)
}

// TODO: implement git bash on windows for type
func HandleType(tokens []string){
	for _, cmd := range tokens[1:] {
		if _, exists := cmds[cmd]; exists {
			fmt.Println(cmd, "is a built-in command")
		} else if lookUpPath, err := exec.LookPath(cmd); err == nil {
			fmt.Println(cmd, "is", lookUpPath)
		} else {
			fmt.Println(cmd, ": invalid command")
		}
	}
}

func HandleWorkingDir(){
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd: encountered error: %v\n", err)
		return
	}

	fmt.Println(pwd)
}

func RunExternalCmd(tokens []string){
	if len(tokens) == 0 {
		return
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		if path, found := utils.HasBash(); found {
			// uses git bash if installed
			cmd = exec.Command(path, "-c", strings.Join(tokens, " "))
		} else {
			cmd = exec.Command("cmd", "/C", strings.Join(tokens, " "))
		}
	} else {
		cmd = exec.Command("sh", "-c", strings.Join(tokens, " "))
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error executing command: %v\n", tokens[0], err)
	}
}

