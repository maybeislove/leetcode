package main

import "fmt"

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		ints := []int{}
		for j := 0; j <= i; j++ {
			ints = append(ints, combi(i, j))
		}
		result = append(result, ints)
	}

	fmt.Println(result)
	return result
}

func combi(r, n int) int {
	p := 1
	for i := 1; i <= n; i++ {
		p = p * (r - i + 1) / i
	}
	return p
}
