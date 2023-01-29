package main

import "fmt"

func main() {
	fmt.Printf("%v", finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Printf("%v", finalPrices([]int{1, 2, 3, 4, 5}))

}

func finalPrices(prices []int) []int {
	result := []int{}
	for i, price := range prices {
		for j := i + 1; j < len(prices); j++ {
			if price >= prices[j] {
				price -= prices[j]
				break
			}
		}
		result = append(result, price)
	}
	return result
}
