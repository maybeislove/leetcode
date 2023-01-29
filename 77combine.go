package main

import "fmt"

func main() {
	fmt.Printf("%v", combine1(4, 2))
}

func combine(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTracking func(n, k, startIndex int)
	backTracking = func(n, k, startIndex int) {
		if len(path) == k {
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[0 : len(path)-1]
		}
	}
	backTracking(n, k, 1)
	return result
}

func combine1(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTracking func(startIndex int)
	backTracking = func(startIndex int) {
		if len(path) == k {
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		//path.size() ： 已经找的个数
		//k-path.size() ：还需找的个数
		//【x, n】的数组长度起码应该是k-path.size()才有继续搜索的可能，
		//那么就有 n-x+1 = k-path.size()，
		//解方程得 x = n+1 - (k-path.size()),
		//而且这个x是可以作为起点往下搜的
		//也就是for(i = s; i<=x; i++) 这里的x是可以取到的
		for i := startIndex; i <= n+1-(k-len(path)); i++ {
			path = append(path, i)
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(1)
	return result
}

func combine2(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
