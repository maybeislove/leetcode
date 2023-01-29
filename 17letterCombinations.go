package main

import "fmt"

func main() {
	fmt.Printf("%v", letterCombinations("23"))
}

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{}
	var backtrack func(digits string, index int, combination string)
	backtrack = func(digits string, index int, combination string) {
		if index == len(digits) {
			result = append(result, combination)
			return
		}
		digit := string(digits[index])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
	backtrack(digits, 0, "")
	return result
}
