package main

func main() {
	println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
	println(minimumRounds([]int{2, 3, 3}))

}

func minimumRounds(tasks []int) int {
	m := make(map[int]int, len(tasks))
	for _, task := range tasks {
		m[task]++
	}
	result := 0

	for _, v := range m {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			result += v / 3
		} else {
			result += v / 3
			result++
		}
	}
	return result
}
