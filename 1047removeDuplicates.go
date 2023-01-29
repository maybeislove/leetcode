package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates1047("abbaca"))
}

func removeDuplicates1047(s string) string {
	result := []byte{}
	for i := range s {
		n := len(result)
		if n > 0 && result[n-1] == s[i] {
			result = result[:n-1]
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
