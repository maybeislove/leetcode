package main

func main() {
	println(fib(5))
}

func fib(n int) int {
	dp := make([]int, n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp[0], dp[1] = 1, 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
