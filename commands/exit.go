package commands

import (
	"fmt"
	"os"

	"github.com/swayamduhan/shell-go/utils"
)

func HandleExit(tokens []utils.Token) {
	if len(tokens) > 2 {
		fmt.Println("too many arguments")
		return
	}
	fmt.Fprint(os.Stdin, "exitted shell\n")
	os.Exit(0)
}