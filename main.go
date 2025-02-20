package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/swayamduhan/shell-go/builtins"
)



func handleCmd(cmd string){
	tokens := strings.Split(cmd, " ")
	
	switch tokens[0] {
	case "echo":
		builtins.HandleEcho(tokens)

	case "exit":
		builtins.HandleExit(tokens)
	case "type":
		builtins.HandleType(tokens)
	default: 
		fmt.Fprint(os.Stderr, cmd + ": command not found\n")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("$> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading command: ", err)
			os.Exit(-1)
		}

		cmd = strings.TrimSpace(cmd)
		handleCmd(cmd)
	}
}