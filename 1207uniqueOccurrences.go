package main

func main() {
	println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	println(uniqueOccurrences([]int{1, 2}))
}

func uniqueOccurrences(arr []int) bool {

	m1, m2 := make(map[int]int, len(arr)), make(map[int]int, len(arr))
	for _, e := range arr {
		m1[e]++
	}
	for _, v := range m1 {
		if val, ok := m2[v]; ok && val > 0 {
			return false
		}
		m2[v]++
	}
	return true
}
