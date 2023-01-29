package main

import "fmt"

func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 9))
}
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[num], i}
		}
		m[num] = i
	}
	return nil
}
