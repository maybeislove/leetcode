package main

func main() {
	intersect([]int{1, 2, 2, 1}, []int{2, 1})
}

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	m := make(map[int]int, len(nums1))

	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if m[num] > 0 {
			result = append(result, num)
			m[num]--
		}
	}
	return result
}
