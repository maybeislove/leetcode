package main

import "strings"

func main() {

}

func repeatedSubstringPattern(s string) bool {

	s2 := strings.Repeat(s, 2)
	return strings.Contains(s2[1:len(s2)-1], s)
}
