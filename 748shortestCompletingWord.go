package main

import (
	"strings"
	"unicode"
)

func main() {
	println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	m := [26]int{}
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) {
			m[ch-'a']++
		}
	}
	result := ""
	for _, word := range words {
		temp := m
		for _, ch := range word {
			temp[ch-'a']--
		}
		f := true
		for _, i := range temp {
			if i > 0 {
				f = false
				break
			}
		}
		if f && result == "" || f && len(word) < len(result) {
			result = word
		}
	}
	return result
}
