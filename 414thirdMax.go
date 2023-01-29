package main

import "math"

func main() {
	println(thirdMax([]int{2, 2, 3, 1}))
}

func thirdMax(nums []int) int {
	min, mid, max := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > max {
			min = mid
			mid = max
			max = num
		} else if max > num && num > mid {
			min = mid
			mid = num
		} else if mid > num && num > min {
			min = num
		}
	}
	if min == math.MinInt64 {
		return max
	}
	return min
}
