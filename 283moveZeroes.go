package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	//moveZeroes([]int{2, 1})
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		var j int
		if nums[i] == 0 {
			for j = i + 1; j < len(nums) && nums[j] == 0; j++ {
			}
		}
		if j == len(nums) {
			break
		}
		if j > 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		}
	}
	fmt.Printf("%v", nums)
}
