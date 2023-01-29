package main

import (
	"math"
	"sort"
)

func main() {
	println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	println(threeSumClosest([]int{0, 0, 0}, 1))

}
func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	//fmt.Printf("%v", nums)
	result := nums[0] + nums[1] + nums[2]
	for i := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			if i == l || i == r {
				continue
			}
			sum := nums[i] + nums[l] + nums[r]

			if sum > target {
				r--
			} else {
				l++
			}
			newResult := math.Abs(float64(sum - target))
			nowResult := math.Abs(float64(result - target))
			if newResult < nowResult {
				result = sum
			}
		}
	}
	return result
}
