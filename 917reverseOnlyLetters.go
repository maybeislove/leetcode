package main

import "unicode"

func main() {
	println(reverseOnlyLetters("ab-cd"))
	println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}

func reverseOnlyLetters(s string) string {
	res := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) {
			l++
		}
		if !unicode.IsLetter(rune(s[r])) {
			r--
		}
		if unicode.IsLetter(rune(s[l])) && unicode.IsLetter(rune(s[r])) {
			res[l], res[r] = s[r], s[l]
			l++
			r--
		}

	}
	return string(res)
}
