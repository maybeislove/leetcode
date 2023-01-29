package main

import "fmt"

func main() {
	fmt.Printf("%v", combinationSum([]int{2, 3, 5}, 4))
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTrack func(startIndex, sum int)
	backTrack = func(startIndex, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backTrack(i, sum)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0, 0)
	return result
}
