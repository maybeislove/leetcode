package main

func main() {
	println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))

}

func countConsistentStrings(allowed string, words []string) int {
	m := [26]bool{}
	for _, ch := range allowed {
		m[ch-'a'] = true
	}
	check := func(s string) bool {
		for _, ch := range s {
			if !m[ch-'a'] {
				return false
			}
		}
		return true
	}
	result := 0
	for _, word := range words {
		if check(word) {
			result++
		}
	}
	return result
}
