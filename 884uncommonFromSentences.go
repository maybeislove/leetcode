package main

import (
	"strings"
)

func main() {
	uncommonFromSentences("this apple is sweet", "this apple is sour")
	uncommonFromSentences("apple apple", "banana")
	uncommonFromSentences("s z z z s", "s z ejt")
	uncommonFromSentences("fo ly ly", "fo fo etx")
}

func uncommonFromSentences(s1 string, s2 string) []string {
	m := make(map[string]int, len(s1))
	result := []string{}
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")
	for _, word := range words1 {
		m[word]++
	}
	for _, word := range words2 {
		m[word]++
	}

	for k, v := range m {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
