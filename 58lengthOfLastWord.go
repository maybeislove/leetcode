package main

import "strings"

func main() {
	s := "Hello World"
	println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	return len(split[len(split)-1])
}
