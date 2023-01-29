package main

import "fmt"

func main() {
	fmt.Println(game([]int{2, 2, 3}, []int{3, 2, 1}))
}

func game(guess []int, answer []int) int {
	result := 0
	for i, val := range guess {
		if val == answer[i] {
			result++
		}
	}
	return result
}
