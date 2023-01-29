package main

import (
	"sort"
)

func main() {
	findDisappearedNumbers([]int{4, 3, 2, 7, 1, 2, 3, 1})

}

func findDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)
	result := []int{}
	m := map[int]struct{}{}
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}
