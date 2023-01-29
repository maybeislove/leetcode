package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(interpret("G()(al)"))
}
func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}
