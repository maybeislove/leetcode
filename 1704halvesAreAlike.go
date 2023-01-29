package main

func main() {
	println(halvesAreAlike("textbook"))
}

var vowel = []int32{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func halvesAreAlike(s string) bool {
	count := 0
	for i, ch := range s {
		if i < len(s)/2 {
			if isVowel(ch) {
				count++
			}
		} else {
			if isVowel(ch) {
				count--
			}
		}
	}
	return count == 0
}

func isVowel(b int32) bool {
	for _, ch := range vowel {
		if b == ch {
			return true
		}
	}
	return false
}
