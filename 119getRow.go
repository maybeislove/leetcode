package main

import "fmt"

func main() {
	getRow(5)
}

func getRow(rowIndex int) []int {
	ints := []int{}
	for j := 0; j <= rowIndex; j++ {
		ints = append(ints, combi(rowIndex, j))
	}
	fmt.Printf("%v", ints)
	return ints
}
