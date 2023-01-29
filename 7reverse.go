package main

import "math"

func main() {
	println(reverse(-12))
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		result *= 10
		result += x % 10
		x /= 10
	}
	if result >= math.MaxInt32 || result <= math.MinInt32 {
		result = 0
	}
	return result
}
