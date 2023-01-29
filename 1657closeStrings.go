package main

import "strings"

func main() {
	println(closeStrings("aav", "vva"))
	println(closeStrings("uau", "ssx"))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	for _, ch := range word1 {
		if !strings.ContainsRune(word2, ch) {
			return false
		}
	}
	m1, m2 := make(map[int32]int, 26), make(map[int32]int, 26)

	for _, ch := range word1 {
		m1[ch-'a']++
	}
	for _, ch := range word2 {
		m2[ch-'a']++
	}

	for k, v := range m1 {
		f := false
		for k2, v2 := range m2 {
			if v == v2 {
				f = true
				delete(m1, k)
				delete(m2, k2)
				break
			}
		}
		if !f {
			return false
		}
	}
	return true
}
