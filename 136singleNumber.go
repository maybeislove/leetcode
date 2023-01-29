package main

func main() {
	number := singleNumber([]int{3, 5, 6, 8, 6, 4, 3, 4, 5})
	println(number)
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
