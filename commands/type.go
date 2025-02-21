package commands

import (
	"fmt"
	"os/exec"
)


var cmds = map[string]bool{
	"echo" : true,
	"exit" : true,
	"type" : true,
	"pwd" : true,
	"cd" : true,
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