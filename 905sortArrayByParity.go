package main

import "fmt"

func main() {
	fmt.Printf("%v", sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]%2 == 0 {
			l++
		} else {
			if nums[r]%2 != 0 {
				r--
			} else {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
		}
	}
	return nums
}
