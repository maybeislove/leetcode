package main

import "fmt"

func main() {
	fmt.Printf("%v", decode([]int{1, 2, 3}, 1))
	fmt.Printf("%v", decode([]int{6, 2, 7, 3}, 4))
}

// first^X = encode[i]
// first^first^X = encode[i]^first
// X= encode[i]^first
func decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)
	result[0] = first

	for i := 1; i < len(result); i++ {
		result[i] = encoded[i-1] ^ result[i-1]
	}
	return result
}
