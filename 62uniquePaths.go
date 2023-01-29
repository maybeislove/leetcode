package main

import "fmt"

func main() {
	println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Printf("%v", dp)
	return dp[m-1][n-1]
}
