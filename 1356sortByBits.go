package main

import (
	"sort"
)

func main() {
	sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		iValue, jValue := arr[i], arr[j]
		iBitNum, jBitNum := countBit1(iValue), countBit1(jValue)
		return iBitNum < jBitNum || iBitNum == jBitNum && iValue < jValue
	})
	return arr
}

func countBit1(n int) int {
	res := 0
	for n != 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
