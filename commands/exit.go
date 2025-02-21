package commands

import (
	"fmt"
	"os"
)

func HandleExit(tokens []string) {
	if len(tokens) > 2 {
		fmt.Println("too many arguments")
		return
	}
	fmt.Fprint(os.Stdin, "exitted shell\n")
	os.Exit(0)
}