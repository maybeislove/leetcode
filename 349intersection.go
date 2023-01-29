package main

func main() {
	intersection2([]int{1, 2, 2, 1}, []int{2, 2})
}
func intersection(nums1 []int, nums2 []int) []int {

	result := []int{}
	m := make(map[int]struct{}, len(nums1))
	resultMap := make(map[int]struct{}, len(nums1))

	for _, num := range nums1 {
		m[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			resultMap[num] = struct{}{}
		}
	}
	for key, _ := range resultMap {
		result = append(result, key)
	}

	//fmt.Printf("%v", result)
	return result
}

func intersection2(nums1 []int, nums2 []int) []int {
	result := []int{}
	m := make(map[int]int, len(nums1))

	for _, num := range nums1 {
		m[num] = 1
	}
	for _, num := range nums2 {
		if val, ok := m[num]; ok && val > 0 {
			result = append(result, num)
			m[num]--
		}
	}
	return result
}
