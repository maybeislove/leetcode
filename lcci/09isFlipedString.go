package main

import "strings"

func main() {
	println(isFlipedString("waterbottle", "erbottlewat"))
}

func isFlipedString(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	var builder strings.Builder
	builder.WriteString(s2)
	builder.WriteString(s2)
	return strings.Contains(builder.String(), s1)
}
