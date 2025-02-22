package commands

import (
	"fmt"

	"github.com/swayamduhan/shell-go/utils"
)


func HandleEcho(tokens []utils.Token) {
	for _, token := range tokens[1:] {
		fmt.Printf("%s ", token.Value)
	}
	fmt.Println()
}