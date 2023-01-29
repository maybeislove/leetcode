package main

import (
	"sort"
)

func main() {
	//println(largestPerimeter([]int{1000, 2, 2, 3}))
	println(largestPerimeter([]int{3, 2, 3, 4}))

}
func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) <= 2 {
		return 0
	}
	//fmt.Printf("%d", nums)
	res := 0
	for i := 2; i < len(nums); i++ {
		if !(nums[i]+nums[i-1] > nums[i-2]) {
			continue
		}
		res = nums[i-1] + nums[i-2] + nums[i]
		break
	}
	return res
}
