package main

import "sort"

func main() {
	println(smallestRangeI([]int{0, 10}, 2))
}
func smallestRangeI(nums []int, k int) int {
	sort.Ints(nums)
	min, max := nums[0], nums[len(nums)-1]
	if max-k > min+k {
		return max - min - 2*k
	}
	return 0
}
