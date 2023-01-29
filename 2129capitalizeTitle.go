package main

import (
	"strings"
)

func main() {
	println(capitalizeTitle("capiTalIze tHe titLe"))
	println(capitalizeTitle("First leTTeR of EACH Word"))

}

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	for i, word := range words {
		if len(word) <= 2 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}
