package commands

import (
	"fmt"
	"os/exec"
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

// TODO: implement git bash on windows for type
func HandleType(tokens []utils.Token){
	for _, cmd := range tokens[1:] {
		if _, exists := cmds[cmd.Value]; exists {
			fmt.Println(cmd.Value, "is a built-in command")
		} else if lookUpPath, err := exec.LookPath(cmd.Value); err == nil {
			fmt.Println(cmd.Value, "is", lookUpPath)
		} else {
			fmt.Println(cmd.Value, ": invalid command")
		}
	}
}

func IsBuiltin(cmd string) bool {
	tokens := strings.Split(cmd, " ")
	if len(tokens) == 0 {
		return false
	}

	_, exists := cmds[tokens[0]]
	return exists
}