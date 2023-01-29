package main

import "fmt"

func main() {
	fmt.Printf("%v", removeElement([]int{2, 3, 1, 3}, 3))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
