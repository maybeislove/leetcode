package main

import "sort"

func main() {
	println(findContentChildren([]int{2, 3, 4}, []int{1, 2, 32}))
}

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)
	res := 0
	var l, r = len(g), len(s)
	for i, j := 0, 0; i < l && j < r; i++ {
		for j < r && i < l {
			if g[i] <= s[j] {
				res++
				j++
				break
			}
			j++
		}
	}

	return res
}
