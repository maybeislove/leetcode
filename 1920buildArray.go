package main

import "fmt"

func main() {
	fmt.Printf("%v", buildArray([]int{0, 2, 1, 5, 3, 4}))
}

func buildArray(nums []int) []int {
	result := make([]int, len(nums))

	for i, num := range nums {
		result[i] = nums[num]
	}
	return result
}
