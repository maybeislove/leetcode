package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseVowels("helloo")
}

func reverseVowels(s string) string {
	b := []byte(s)
	var left, right = 0, len(s) - 1

	for left < right {
		if !isVowels(s[left]) {
			left++
		}
		if !isVowels(s[right]) {
			right--
		}
		if isVowels(s[left]) && isVowels(s[right]) {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}
	fmt.Printf("%v", string(b))
	return string(b)
}

func isVowels(u uint8) bool {
	return strings.Contains("aeiouAEIOU", string(u))
}
