package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v", findOcurrences("alice is a good girl she is a good student", "a", "good"))

}

func findOcurrences(text string, first string, second string) []string {
	result := []string{}
	words := strings.Split(text, " ")

	for i := 2; i < len(words); i++ {
		if words[i-1] == second && words[i-2] == first {
			result = append(result, words[i])
		}
	}
	return result
}
