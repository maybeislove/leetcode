package main

func main() {
	println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}
func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int, len(target))

	for i, element := range target {
		m[element]++
		m[arr[i]]--
	}

	for _, element := range m {
		if element != 0 {
			return false
		}
	}
	return true
}
