package main

func main() {
	//println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))

}

func findMaxConsecutiveOnes(nums []int) int {
	count, result := 0, 0

	for _, num := range nums {
		if num == 1 {
			count++
		}
		if count > result {
			result = count
		}
		if num == 0 {
			count = 0
		}
	}
	return result
}
