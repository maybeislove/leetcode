package main

import (
	"sort"
)

func main() {
	//println(isMonotonic([]int{1, 2, 2, 3}))
	//println(isMonotonic([]int{1, 3, 2}))
	//println(isMonotonic([]int{1, 1, 2}))
	println(isMonotonic([]int{1, 1, 0}))

}

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	a := make([]int, len(nums))
	if nums[0]-nums[1] == 0 {
		return isMonotonic(nums[1:])
	}
	copy(a, nums)
	if nums[0]-nums[1] <= 0 {
		sort.Ints(nums)
		return isEqual(a, nums)
	} else {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})
		return isEqual(a, nums)
	}
}

func isEqual(a, b []int) bool {
	for i, e := range a {
		if b[i] != e {
			return false
		}
	}
	return true
}

func isMonotonic1(nums []int) bool {
	return sort.IntsAreSorted(nums) || sort.IsSorted(sort.Reverse(sort.IntSlice(nums)))
}
