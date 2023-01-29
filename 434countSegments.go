package main

import (
	"strings"
)

func main() {
	println(countSegments("hello worlkd"))
}

func countSegments(s string) int {
	res := 0
	ss := strings.Split(s, " ")

	for _, ch := range ss {
		if ch != "" {
			res++
		}
	}
	return res
}
