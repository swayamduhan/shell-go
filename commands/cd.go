package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/swayamduhan/shell-go/utils"
)

var lastDir string // to store previous directory for "cd -"

// add a additional command using DLL for surfing forward and previous like a browser
func HandleChangeDir(tokens []string) {
	if len(tokens) > 2 {
		fmt.Fprintf(os.Stderr, "cd: too many arguments\n")
		return
	}

	var path string
	if len(tokens) == 1 || tokens[1] == "~" {
		path, _ = utils.GetHomeDir()
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