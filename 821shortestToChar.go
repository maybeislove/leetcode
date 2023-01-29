package main

import "fmt"

func main() {
	shortestToChar("loveleetcode", 'e')
}

func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))
	idx := -len(s)
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		result[i] = i - idx
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		result[i] = min821(idx-i, result[i])
	}
	fmt.Printf("%v", result)

	return result
}

func min821(a, b int) int {
	if a < b {
		return a
	}
	return b
}
