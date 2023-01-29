package main

import (
	"sort"
)

func main() {
	merge([]int{1}, 1, []int{}, 0)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]
	if n != 0 {
		nums1 = append(nums1, nums2[0:n]...)
	}
	sort.Ints(nums1)
}
