package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func replaceElements(arr []int) []int {
	result := []int{}
	for i, _ := range arr {
		max := math.MinInt32
		for j := i + 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
			}
		}
		result = append(result, max)
	}
	result[len(arr)-1] = -1
	return result
}
