package main

import (
	"strings"
)

func main() {
	println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(sentence string) string {
	result := ""
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.Contains("aeiouAEIOU", string(word[0])) {
			word += "ma"
		} else {
			word += string(word[0]) + "ma"
			word = word[1:]
		}

		for j := 0; j <= i; j++ {
			word += "a"
		}
		words[i] = word
		result = strings.Join(words, " ")
	}
	return result

}
