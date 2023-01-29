package main

import "math"

func main() {
	println(isHappy(2))
	//println(nextN(16))
}

func isHappy(n int) bool {
	m := make(map[int]int)

	for n >= 2 {
		m[n] = n
		n = nextN(n)
		if _, ok := m[n]; ok {
			return false
		}
	}
	return true
}

func nextN(n int) int {
	result := 0
	for n > 0 {
		result += int(math.Pow(float64(n%10), 2))
		n /= 10
	}
	return result
}
