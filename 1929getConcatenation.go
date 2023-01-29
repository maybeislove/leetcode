package main

import "fmt"

func main() {
	fmt.Printf("%v", getConcatenation([]int{1, 3, 2, 1}))
}

func getConcatenation(nums []int) []int {
	for _, num := range nums {
		nums = append(nums, num)
	}
	return nums
}
