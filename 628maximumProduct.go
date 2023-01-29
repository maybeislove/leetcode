package main

import (
	"math"
	"sort"
)

func main() {
	println(maximumProduct([]int{10, 1, 2, 3}))

}
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return int(math.Max(float64(nums[0]*nums[1]*nums[len(nums)-1]), float64(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])))
}
