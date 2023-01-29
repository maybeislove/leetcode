package main

import "fmt"

func main() {
	println(rob([]int{1, 2, 3, 1}))
	println(rob([]int{2, 7, 9, 3, 1}))

}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	fmt.Printf("%v", dp)
	return dp[len(nums)-1]
}

func Max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
