package main

import "fmt"

func main() {

	fmt.Printf("%v", sumZero(5))
}

func sumZero(n int) []int {
	if n == 0 || n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{-1, 1}
	}
	ints := make([]int, n)

	for i := 0; i < n-1; i++ {
		ints[i] = i
	}

	ints[n-1] = -1 * (1 + n - 2) * (n - 2) / 2
	return ints
}
