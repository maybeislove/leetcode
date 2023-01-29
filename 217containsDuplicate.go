package main

func main() {
	println(containsDuplicate([]int{1, 2, 3, 4}))
}

func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
