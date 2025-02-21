package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/swayamduhan/shell-go/utils"
)

var cmds = map[string]bool{
	"echo" : true,
	"exit" : true,
	"type" : true,
	"pwd" : true,
	"cd" : true,
}

var lastDir string  // to store previous directory for "cd -"

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

func HandleChangeDir(tokens []string){
	if len(tokens) > 2 {
		fmt.Fprintf(os.Stderr, "cd: too many arguments\n")
		return
	}

	var path string
	if len(tokens) == 1 || tokens[1] == "~" {
		path = os.Getenv("HOME")
	} else if tokens[1] == "-" {
		if lastDir == "" {
			fmt.Fprintf(os.Stderr, "cd: OLDPWD not set\n")
			return
		}
		path = lastDir
	} else {
		path = tokens[1]
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: error resolving path: %v\n", err)
		return
	}

	currDir, _ := os.Getwd()
	lastDir = currDir


	err = os.Chdir(absPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: cannot access '%s': No such file or directory\n", path)
	}
}

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

