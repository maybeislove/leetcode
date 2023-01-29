package main

import "fmt"

func main() {
	println(climbStairs(4))
	println(climbStairs3(4))

}

func climbStairs(n int) int {
	var p, q, r = 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r

}
func climbStairs2(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Printf("%v", dp)
	return dp[n]

}

func climbStairs3(n int) int {
	var dp = make([]int, 3)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		sum := dp[1] + dp[2]
		dp[1] = dp[2]
		dp[2] = sum
	}
	fmt.Printf("%v", dp)
	return dp[2]

}
