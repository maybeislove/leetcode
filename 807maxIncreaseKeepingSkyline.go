package main

import (
	"math"
)

func main() {
	println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax, colMax := make([]int, len(grid)), make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			rowMax[i] = int(math.Max(float64(rowMax[i]), float64(grid[i][j])))
			colMax[j] = int(math.Max(float64(colMax[j]), float64(grid[i][j])))
		}
	}
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			result += int(math.Min(float64(rowMax[i]), float64(colMax[j]))) - grid[i][j]
		}
	}
	return result
}
