package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", minOperations1769("110"))
	fmt.Printf("%v", minOperations1769("001011"))

}

func minOperations1769(boxes string) []int {
	result := make([]int, len(boxes))
	for i := range boxes {
		temp := 0
		for j := range boxes {
			if i == j {
				continue
			}
			if boxes[j] == '1' {
				temp += int(math.Abs(float64(i - j)))
			}
		}
		result[i] = temp
	}
	return result
}
