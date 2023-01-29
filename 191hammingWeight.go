package main

import "fmt"

func main() {
	hammingWeight(7)
}

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		if num%2 == 1 {
			result += 1
		}
		num = num / 2
	}
	fmt.Println(result)
	return result
}
