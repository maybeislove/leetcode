package main

func main() {
	println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	println(findKthPositive([]int{1, 2, 3, 4}, 2))

}
func findKthPositive(arr []int, k int) int {
	m := make(map[int]struct{}, len(arr))
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = struct{}{}
	}
	for i := 1; i <= 2000; i++ {
		if _, ok := m[i]; !ok {
			k--
			if k == 0 {
				return i
			}
		}
	}
	return 0
}
