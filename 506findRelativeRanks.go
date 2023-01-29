package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("%v", findRelativeRanks([]int{10, 3, 8, 9, 4}))

}

func findRelativeRanks(score []int) []string {
	result := []string{}
	revNums := make([]int, len(score))
	copy(revNums, score)
	sort.Slice(revNums, func(i, j int) bool {
		return revNums[i] > revNums[j]
	})
	m := make(map[int]int, len(revNums))
	for i, num := range revNums {
		m[num] = i + 1
	}
	for _, s := range score {
		if m[s] == 1 {
			result = append(result, "Gold Medal")
		} else if m[s] == 2 {
			result = append(result, "Silver Medal")
		} else if m[s] == 3 {
			result = append(result, "Bronze Medal")
		} else {
			result = append(result, strconv.Itoa(m[s]))
		}
	}
	return result
}
