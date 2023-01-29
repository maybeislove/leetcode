package main

import (
	"sort"
)

func main() {
	println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	result := 0
	for i := 0; i < len(rocks); i++ {
		capacity[i] -= rocks[i]
	}
	sort.Ints(capacity)
	for _, d := range capacity {
		if additionalRocks < d {
			break
		}
		additionalRocks -= d
		result++

	}
	return result
}
