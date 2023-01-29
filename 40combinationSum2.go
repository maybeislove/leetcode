package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}
	sort.Ints(candidates)
	var backTracking func(startIndex, target, sum int)
	backTracking = func(startIndex, target, sum int) {
		if sum > target {
			return
		}
		if target == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			if i > startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backTracking(i+1, target, sum)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTracking(0, target, 0)
	return result
}
