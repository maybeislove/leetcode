package main

import "strings"

func main() {
	println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {

	words := strings.Split(s, " ")
	revwords := []string{}
	for _, word := range words {
		byte_str := []rune(word)
		for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
			byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
		}
		revwords = append(revwords, string(byte_str))
	}
	return strings.Join(revwords, " ")
}
