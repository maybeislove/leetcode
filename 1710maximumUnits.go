package main

import (
	"sort"
)

func main() {
	println(maximumUnits([][]int{{1, 3}, {2, 6}, {3, 2}}, 4))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	result := 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, boxType := range boxTypes {
		if truckSize > 0 {
			if truckSize >= boxType[0] {
				truckSize -= boxType[0]
				result += boxType[0] * boxType[1]
			} else {
				result += truckSize * boxType[1]
				break
			}
		}
	}
	return result
}
