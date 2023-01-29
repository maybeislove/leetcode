package main

import "fmt"

func main() {
	fmt.Printf("%v", countBits(5))
}

func countBits(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, countBit(i))
	}
	return result
}

func countBit(n int) int {
	res := 0
	for n != 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
