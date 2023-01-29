package main

import "fmt"

func main() {
	fmt.Printf("%v", shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}

func shuffle(nums []int, n int) []int {
	result := make([]int, len(nums))
	for i := 0; i < n; i++ {
		result[i*2] = nums[i]
		result[2*i+1] = nums[i+n]
	}
	return result
}
