package main

import "fmt"

func main() {
	fmt.Printf("%v", nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Printf("%v", nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))

}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := []int{}
	m := map[int]int{}
	for i, num2 := range nums2 {
		if i == len(nums2)-1 {
			break
		}
		j := i + 1
		for ; j < len(nums2)-1; j++ {
			if nums2[j] > nums2[i] {
				break
			}
		}
		m[num2] = nums2[j]
	}
	for _, num1 := range nums1 {
		if val, ok := m[num1]; ok && val > num1 {
			result = append(result, val)
		} else {
			result = append(result, -1)
		}
	}
	return result
}
