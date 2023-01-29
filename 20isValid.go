package main

import (
	"strings"
)

func main() {
	println(isValid("(([]){})"))
}

func isValid(s string) bool {
	l := len(s) / 2
	for i := 0; i < l; i++ {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
		if len(s) == 0 {
			break
		}
	}
	return len(s) == 0
}
