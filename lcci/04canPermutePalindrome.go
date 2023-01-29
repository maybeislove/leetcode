package main

func main() {
	println(canPermutePalindrome("aaasss"))
}

func canPermutePalindrome(s string) bool {

	m := make(map[int32]int, len(s))

	for _, ch := range s {
		m[ch]++
	}
	count := 0
	for _, val := range m {
		if val%2 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
