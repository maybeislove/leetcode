package main

import "fmt"

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}

func checkIfPangram(sentence string) bool {
	s := [26]int32{}
	for _, ch := range sentence {
		s[ch-'a']++
	}
	for _, val := range s {
		if val == 0 {
			return false
		}
	}
	return true
}
