package main

import "sort"

func main() {
	println(dominantIndex([]int{3, 6, 1, 0}))
}

func dominantIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	sort.Ints(nums)
	maxIndex := len(nums) - 1
	if nums[maxIndex] >= nums[maxIndex-1]*2 {
		return m[nums[len(nums)-1]]
	}
	return -1
}
