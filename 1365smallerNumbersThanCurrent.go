package main

import "fmt"

func main() {
	fmt.Printf("%v", smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i := range nums {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] > nums[j] {
				result[i]++
			}
		}
	}
	return result
}
