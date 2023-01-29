package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	var result []string
	for i, n := 0, len(nums); i < n; {
		var left = i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if i-1 > left {
			s += fmt.Sprintf("->%d", nums[i-1])
		}
		result = append(result, s)
	}
	return result
}
