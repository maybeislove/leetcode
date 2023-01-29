package main

import "strings"

func main() {
	println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))

}

func mostWordsFound(sentences []string) int {
	result := 0
	for _, sentence := range sentences {
		s := strings.Split(sentence, " ")
		if len(s) > result {
			result = len(s)
		}
	}
	return result
}
