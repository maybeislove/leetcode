package main

import (
	"sort"
)

func main() {
	lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
}

func lastStoneWeight(stones []int) int {

	compare := func(stones []int) []int {
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
		i := stones[0] - stones[1]
		if i != 0 {
			stones = stones[2:]
			stones = append(stones, i)
			return stones
		}
		stones = stones[2:]
		return stones
	}

	for len(stones) > 1 {
		stones = compare(stones)
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
