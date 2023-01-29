package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v", findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Printf("%v", findWords([]string{"a", "b"}))
}

func findWords(words []string) []string {
	result := []string{}
	for i, word := range words {
		row := whichRow(words[i])
		flag := true
		for _, ch := range word {
			if !strings.Contains(row, string(ch)) {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, word)
		}
	}
	return result
}

func whichRow(word string) string {
	if strings.ContainsAny(row1, word[0:1]) {
		return row1
	}
	if strings.ContainsAny(row2, word[0:1]) {
		return row2
	}
	if strings.ContainsAny(row3, word[0:1]) {
		return row3
	}
	return ""
}

const (
	row1 string = "QWERTYUIOPqwertyuiop"
	row2 string = "ASDFGHJKLasdfghjkl"
	row3 string = "ZXCVBNMzxcvbnm"
)
