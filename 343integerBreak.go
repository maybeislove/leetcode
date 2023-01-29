package main

import "fmt"

func main() {
	integerBreak(10)
}

func integerBreak(n int) int {

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1
	for i := 2; i <= n; i++ {
		temp := 0
		for j := 1; j <= i/2; j++ {
			temp = max(temp, max(j*dp[i-j], j*(i-j)))
		}
		dp[i] = temp
	}
	fmt.Printf("%v", dp)
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
