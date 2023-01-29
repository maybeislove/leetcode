package main

func main() {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

func duplicateZeros(arr []int) {
	n := len(arr)
	top := 0
	i := -1
	for top < n {
		i++
		if arr[i] != 0 {
			top++
		} else {
			top += 2
		}
	}
	j := n - 1
	if top == n+1 {
		arr[j] = 0
		j--
		i--
	}
	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = arr[i]
			j--
		}
		i--
	}

}
