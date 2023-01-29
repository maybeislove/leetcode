package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSum(2932))
}

func minimumSum(num int) int {
	ints := make([]int, 4)
	for i := 0; i < 4; i++ {
		ints[i] = num % 10
		num /= 10
	}
	sort.Ints(ints)
	return 10*ints[0] + ints[2] + 10*ints[1] + ints[3]
}
