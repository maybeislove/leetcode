package main

func main() {
	println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
}
func minCostToMoveChips(position []int) int {
	odd, even, res := 0, 0, 0

	for i := 0; i < len(position); i++ {
		if position[i]%2 != 0 {
			odd++
		} else {
			even++
		}
	}
	if odd > even {
		res = even
		return res
	}
	return odd
}
