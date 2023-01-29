package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", minSubsequence([]int{4, 3, 10, 9, 8}))
	fmt.Printf("%v", minSubsequence([]int{4, 4, 7, 6, 7}))
}
func minSubsequence(nums []int) []int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	result := []int{}
	count := 0
	for _, num := range nums {
		total -= num
		count += num
		result = append(result, num)
		if count > total {
			break
		}

	}
	return result
}
