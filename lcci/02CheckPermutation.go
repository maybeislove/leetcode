package main

import (
	"sort"
)

func main() {
	CheckPermutation2("favd", "")
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[int32]int, len(s1))
	for _, ch := range s1 {
		m[ch]++
	}

	for _, ch := range s2 {
		if val, ok := m[ch]; !ok || val <= 0 {
			return false
		}
		m[ch]--
	}
	return true
}
func CheckPermutation2(s1 string, s2 string) bool {
	words1 := []byte(s1)
	sort.Slice(words1, func(i, j int) bool {
		return words1[i] < words1[j]
	})

	words2 := []byte(s2)
	sort.Slice(words2, func(i, j int) bool {
		return words2[i] < words2[j]
	})
	return string(words1) == string(words2)
}
