package main

import "sort"

func main() {
	println(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
