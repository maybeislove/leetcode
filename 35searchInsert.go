package main

func main() {
	println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)-1
	result := len(nums)

	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] >= target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}
