package main

func main() {
	//println(findLengthOfLCIS([]int{1, 3, 5, 7, 9}))
	//println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
	//println(findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5}))
	println(findLengthOfLCIS([]int{1, 2, 5, 8, 0}))

}

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	result := 0

	count := 0
	for i := 1; i < len(nums); i++ {
		r := nums[i] - nums[i-1]
		if r <= 0 {
			count = 0
			continue
		}
		count++
		if count > result {
			result = count
		}
	}
	return result + 1
}
