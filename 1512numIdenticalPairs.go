package main

func main() {
	println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}

func numIdenticalPairs(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result++
			}
		}
	}
	return result
}
