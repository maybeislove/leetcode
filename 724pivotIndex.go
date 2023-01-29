package main

func main() {
	println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) //28
	println(pivotIndex([]int{1, 2, 3}))
	println(pivotIndex([]int{100, 2, -2}))
}

//左求和*2+中心索引值 = 总和
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		if 2*sum+num == total {
			return i
		}
		sum += num
	}
	return -1
}
