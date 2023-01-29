package main

func main() {
	println(missingNumber([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	//math
	i := (1 + len(nums)) * len(nums) / 2
	for _, num := range nums {
		i -= num
	}
	return i
}
