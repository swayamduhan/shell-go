package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/swayamduhan/shell-go/commands"
)

const (
	// ANSI color codes
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Bold   = "\033[1m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
)



func handleCmd(cmd string){
	tokens := strings.Split(cmd, " ")
	
	switch tokens[0] {
	case "echo":
		commands.HandleEcho(tokens)
	case "exit":
		commands.HandleExit(tokens)
	case "type":
		commands.HandleType(tokens)
	case "pwd":
		commands.HandleWorkingDir()
	case "cd":
		commands.HandleChangeDir(tokens)
	case "":
		fmt.Println("enter a command")
	default: 
		commands.RunExternalCmd(tokens)
	}
}



// implement readline for handling autocompletion and syntax hightlighting
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		path, _ := os.Getwd()
		fmt.Printf("%s%s$%s %s%s%s> ", Bold, Green, Reset, Cyan, path, Reset)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading command: ", err)
			os.Exit(-1)
		}

		cmd = strings.TrimSpace(cmd)
		handleCmd(cmd)
	}

}