package main

func main() {
	println(maxProfit([]int{7, 2, 3, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var min = prices[0]
	var profit = 0

	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		tempProfit := prices[i] - min
		if tempProfit > 0 && tempProfit > profit {
			profit = tempProfit
		}
	}

	return profit
}
