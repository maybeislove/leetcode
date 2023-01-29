package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
}

func reverseLeftWords(s string, n int) string {
	m := len(s)
	s += s
	return s[n : m+n]
}
