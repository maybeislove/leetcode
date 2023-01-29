package main

import (
	"math"
	"sort"
)

func main() {
	println(findLHS([]int{1, 3, 4, 6, 7, 7, 7, 7}))

}

func findLHS(nums []int) int {
	sort.Ints(nums)
	var l, result = 0, 0
	for r, num := range nums {
		for num-nums[l] > 1 {
			l++
		}
		if nums[r]-nums[l] == 1 {
			result = int(math.Max(float64(r-l+1), float64(result)))
		}
	}
	return result
}
