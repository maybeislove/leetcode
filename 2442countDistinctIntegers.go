package main

import "fmt"

func main() {
	fmt.Printf("%v", countDistinctIntegers([]int{1, 13, 10, 12, 31}))
}

func countDistinctIntegers(nums []int) int {
	for _, num := range nums {
		nums = append(nums, revert(num))
	}
	m := make(map[int]struct{}, 2*len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	return len(m)
}

func revert(num int) int {
	res := 0
	for num > 0 {
		mod := num % 10
		res = res*10 + mod
		num /= 10
	}
	return res
}
