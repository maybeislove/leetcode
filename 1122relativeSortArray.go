package main

import (
	"fmt"
	"sort"
)

func main() {
	//array := relativeSortArray([]int{19, 2, 3, 1, 3, 2, 4, 6, 7, 9, 2}, []int{2, 1, 4, 3, 9, 6})
	array := relativeSortArray([]int{26, 21, 11, 20, 50, 34, 1, 18}, []int{21, 11, 26, 20})
	fmt.Printf("%v", array)

}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := map[int]int{}

	for i, arr := range arr2 {
		m[arr] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		indexX, hasX := m[x]
		indexY, hasY := m[y]
		if hasX && hasY {
			return indexX < indexY
		}
		if hasX || hasY {
			println(hasX)
			println(hasY)

			return hasX
		}
		return x < y
	})
	return arr1
}
