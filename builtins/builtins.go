package builtins

import (
	"fmt"
	"os"
	"strings"
)

var cmds = map[string]bool{
	"echo" : true,
	"exit" : true,
	"type" : true,
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

func HandleType(tokens []string){
	for _, cmd := range tokens[1:] {
		if _, exists := cmds[cmd]; exists {
			fmt.Println(cmd, "is a built-in command")
		} else {
			fmt.Println(cmd, ": invalid command")
		}
	}
}


