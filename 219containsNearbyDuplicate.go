package main

import "math"

func main() {
	println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if val, ok := m[num]; ok && int(math.Abs(float64(val-i))) <= k {
			return true
		}
		m[num] = i
	}
	return false
}
