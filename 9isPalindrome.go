package main

import "strconv"

func main() {
	println(isPalindrome(123))
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	palindromes := strconv.Itoa(x)
	for i, _ := range palindromes {
		if palindromes[i] != palindromes[len(palindromes)-i-1] {
			return false
		}
		if i == len(palindromes)/2 {
			return true
		}
	}
	return false
}
