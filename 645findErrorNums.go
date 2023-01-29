package main

func main() {
	findErrorNums([]int{1, 2, 2, 4})
}

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}

	for i := 1; i <= len(nums); i++ {
		if a := m[i]; a == 2 {
			result[0] = i
		} else if a == 0 {
			result[1] = i
		}
	}
	//fmt.Printf("%v", result)
	return result
}
