package main

func main() {
	println(findNumbers([]int{12, 345, 2, 6, 7896}))
	println(findNumbers([]int{555, 901, 482, 1771}))

}

func findNumbers(nums []int) int {
	result := 0
	for _, num := range nums {
		if isEven(num) {
			result++
		}
	}
	return result
}

func isEven(num int) bool {
	count := 0
	for num > 0 {
		count++
		num /= 10
	}
	return count%2 == 0
}
