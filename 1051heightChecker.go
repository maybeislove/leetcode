package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

func heightChecker(heights []int) int {
	result := 0
	//h := make([]int, len(heights))
	//copy(h, heights)

	h := append([]int{}, heights...)
	sort.Ints(heights)

	for i, height := range heights {
		if height != h[i] {
			result++
		}
	}
	return result
}
