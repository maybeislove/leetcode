package main

func main() {
	println(validMountainArray([]int{0, 3, 2, 1}))
}

func validMountainArray(arr []int) bool {
	max := 0
	maxIndex := 0
	for i, a := range arr {
		if a > max {
			max, maxIndex = a, i
		}
	}
	if maxIndex == 0 || maxIndex == len(arr)-1 {
		return false
	}
	for i := 0; i < maxIndex; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	for i := maxIndex; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}
