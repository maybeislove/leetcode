package main

import "math"

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	ans, l, r := 0, 0, len(height)-1
	for l < r {
		area := int(math.Min(float64(height[l]), float64(height[r]))) * (r - l)
		if area > ans {
			ans = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}
