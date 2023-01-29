package main

import (
	"strconv"
	"strings"
)

func main() {
	println(compressString("aabcccccaaa"))

}

func compressString(S string) string {
	builder := strings.Builder{}
	if len(S) == 0 {
		return ""
	}
	count := 1
	char := S[0]
	for i := 1; i < len(S); i++ {
		if S[i] == char {
			count++
		} else {
			builder.WriteString(string(char) + strconv.Itoa(count))
			count = 1
			char = S[i]
		}
	}
	builder.WriteString(string(char) + strconv.Itoa(count))
	if builder.Len() >= len(S) {
		return S
	}
	return builder.String()
}
