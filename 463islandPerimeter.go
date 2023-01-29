package main

func main() {
	println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
}

func islandPerimeter(grid [][]int) int {
	result := 0
	for i, ints := range grid {
		for j, val := range ints {
			if val == 0 {
				continue
			}
			island := 0
			if j-1 >= 0 && ints[j-1] == 1 {
				island++
			}
			if j+1 < len(ints) && ints[j+1] == 1 {
				island++
			}
			if i-1 >= 0 && grid[i-1][j] == 1 {
				island++
			}
			if i+1 < len(grid) && grid[i+1][j] == 1 {
				island++
			}

			result += 4 - island
		}
	}
	return result
}
