package main

import (
	"math"
)

func main() {
	println(timeRequiredToBuy([]int{2, 3, 2}, 2))
}
func timeRequiredToBuy(tickets []int, k int) int {
	result := 0
	for i, ticket := range tickets {
		if i <= k {
			result += int(math.Min(float64(ticket), float64(tickets[k])))
		} else {
			result += int(math.Min(float64(ticket), float64(tickets[k]-1)))
		}
	}
	return result
}
