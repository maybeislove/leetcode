package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumNums(3))
}
func sumNums(n int) int {
	result := 0
	var sum func(i int) bool
	sum = func(i int) bool {
		result += i
		return i > 0 && sum(i-1)
	}
	sum(n)
	return result
}
