package main

func main() {
	println(canJump([]int{3, 2, 1, 0, 4}))
	println(canJump([]int{2, 3, 1, 1, 4}))

}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	greedy := 0
	for i := 0; i <= greedy; i++ {
		greedy = max1(greedy, i+nums[i])
		if greedy >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
