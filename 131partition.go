package main

import "fmt"

func main() {
	fmt.Printf("%v", partition("aab"))
}

func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}
	var backTracking func(s string, startIndex int)
	backTracking = func(s string, startIndex int) {
		if len(s) == startIndex {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]
			if isPalindromes1(str) {
				path = append(path, str)
				backTracking(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	backTracking(s, 0)
	return result
}

func isPalindromes1(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
