package main

import "math"

func main() {
	println(minStartValue([]int{-3, 2, -3, 4, 2}))
	println(minStartValue([]int{1, 2}))

}

func minStartValue(nums []int) int {
	value, startValue := 0, 1
	for _, num := range nums {
		value += num

		if value < 1 {
			startValue = int(math.Max(float64(startValue), float64(1-value)))
		}
	}
	return startValue
}
