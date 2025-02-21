package commands

import (
	"fmt"
	"strings"
)

func HandleEcho(tokens []string) {
	fmt.Println(strings.Join(tokens[1:], " "))
}