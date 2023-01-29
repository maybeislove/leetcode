package main

import "math"

func main() {

}

func countElements(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max, min := math.MinInt32, math.MaxInt32
	result := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	for _, num := range nums {
		if num != max && num != min {
			result++
		}
	}
	return result
}
