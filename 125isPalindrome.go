package main

import (
	"strings"
)

func main() {
	println(isPalindromes("A man, a plan, a canal: Panama"))
}

func isPalindromes(s string) bool {
	s = strings.ToLower(s)
	alphaBet := "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < len(s); i++ {
		if !strings.Contains(alphaBet, string(s[i])) {
			s = strings.Replace(s, string(s[i]), "", -1)
			i--
		}
	}
	println(s)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		return false
	}
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
