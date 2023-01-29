package main

func main() {
	println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}

func findSpecialInteger(arr []int) int {
	m := make(map[int]int, len(arr))
	max := len(arr) / 4
	for _, element := range arr {
		m[element]++

		if m[element] > max {
			return element
		}
	}
	return 0
}
