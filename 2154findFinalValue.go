package main

func main() {
	println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
}

func findFinalValue(nums []int, original int) int {
	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		m[num] = struct{}{}
	}

	for {
		if _, ok := m[original]; ok {
			original *= 2
			continue
		}
		return original
	}

}
