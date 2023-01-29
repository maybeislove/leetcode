package main

import (
	"fmt"
	"sort"
)

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("%v", threeSum([]int{0, 0, 0, 0}))
	fmt.Printf("%v", threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Printf("%v", threeSum([]int{3, 0, -2, -1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := [][]int{}
	for i, num := range nums {
		if num > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				result = append(result, []int{nums[l], nums[r], nums[i]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}
