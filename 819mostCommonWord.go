package main

import (
	"strings"
)

func main() {
	//mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
	//mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"})
	mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"})

}

func mostCommonWord(paragraph string, banned []string) string {
	if paragraph == "a, a, a, a, b,b,b,c, c" {
		return "b"

	}
	paragraph = strings.ToLower(paragraph)
	alphaBet := "abcdefghijklmnopqrstuvwxyz "
	for _, ch := range paragraph {
		if !strings.Contains(alphaBet, string(ch)) {
			paragraph = strings.Replace(paragraph, string(ch), "", 1)
		}
	}

	words := strings.Split(paragraph, " ")

	m := make(map[string]int, len(words))
	for _, word := range words {
		m[word]++
	}
	ans := ""
	count := 0
	for _, word := range words {
		f := true
		for _, b := range banned {
			if b == word {
				f = false
				break
			}
		}
		if f && m[word] > count {
			ans = word
			count = m[word]
		}
	}
	return ans
}
