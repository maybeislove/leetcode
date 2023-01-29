package main

import "fmt"

func main() {
	fmt.Printf("%v", fairCandySwap([]int{1, 1}, []int{2, 2}))
}
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	m := make(map[int]struct{})
	sumA, sumB := 0, 0
	for _, alice := range aliceSizes {
		sumA += alice
		m[alice] = struct{}{}
	}
	for _, bob := range bobSizes {
		sumB += bob
	}
	mid := (sumA - sumB) / 2
	for i := 0; ; i++ {
		y := bobSizes[i]
		x := y + mid
		if _, ok := m[x]; ok {
			return []int{x, y}
		}
	}
}
