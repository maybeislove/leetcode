package main

import "fmt"

func main() {
	fmt.Printf("%v", combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTrack func(startIndex int)
	backTrack = func(startIndex int) {
		if value(path) > n {
			return
		}
		if len(path) == k {
			if value(path) == n {
				temp := make([]int, k)
				copy(temp, path)
				result = append(result, temp)
			}
			return
		}
		for i := startIndex; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(1)
	return result
}

func value(path []int) int {
	sum := 0
	for _, val := range path {
		sum += val
	}
	return sum
}
