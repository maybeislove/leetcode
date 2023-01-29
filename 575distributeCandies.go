package main

func main() {
	println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	println(distributeCandies([]int{6, 6, 6, 6}))

}

func distributeCandies(candyType []int) int {
	result := 0
	m := make(map[int]struct{}, len(candyType))
	for _, candy := range candyType {
		if _, ok := m[candy]; !ok {
			result++
			m[candy] = struct{}{}
		}

		if result >= len(candyType)/2 {
			return result
		}
	}
	return result
}
