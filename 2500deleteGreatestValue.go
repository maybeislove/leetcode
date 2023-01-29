package main

import (
	"sort"
)

func main() {
	deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}})
}

func deleteGreatestValue(grid [][]int) int {
	result := 0
	for _, g := range grid {
		sort.Ints(g)
	}
	for j := range grid[0] {
		mx := 0
		for _, row := range grid {
			mx = max2(mx, row[j])
		}
		result += mx
	}
	return result
}

func max2(a, b int) int {
	if b > a {
		return b
	}
	return a
}
