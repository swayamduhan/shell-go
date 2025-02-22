package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/swayamduhan/shell-go/commands"
	"github.com/swayamduhan/shell-go/utils"
)

const (
	// ANSI color codes
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Bold   = "\033[1m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
)



func handleBuiltinCmd(tokens []utils.Token){
	
	switch tokens[0].Value {
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
	default:
		fmt.Println("enter a command")
	}
}


// add stack based balancing to check for incomplete commands ( quotations and parenthesis ) and let user complete the line
// implement readline for handling autocompletion and syntax 


// FLOW : 
// 1. read line
// 2. check balancing
// 3. if balanced, then tokenize and send for command
// 4. if unbalanced, then ask again for complete input, append it to previously entered string and send for command if balanced


var balanced = false

func main() {

	// set home directory when terminal starts
	home, err := utils.GetHomeDir()
	if err != nil {
		fmt.Println("error getting home directory!")
		return
	}
	os.Chdir(home)

	reader := bufio.NewReader(os.Stdin)
	for {
		path := utils.GetDir()
		if path == "" {
			path = "home"
		}

		fmt.Printf("%s%s%s> %s%s$%s ", Cyan, path, Reset, Bold, Green, Reset)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading command: ", err)
			os.Exit(-1)
		}

		
		cmd = strings.TrimSpace(cmd)

		if commands.IsBuiltin(cmd) {
			// tokenize and pass in the command
			tokenizedInstruction := utils.Tokenize(cmd)
			handleBuiltinCmd(tokenizedInstruction)
		} else {
			commands.RunExternalCmd(strings.Split(cmd, " "))
		}
	}

}