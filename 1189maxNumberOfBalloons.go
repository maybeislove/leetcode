package main

func main() {
	println(maxNumberOfBalloons("loonbalxballpoon"))
}

func maxNumberOfBalloons(text string) int {
	arr := make([]int, 5)
	for _, ch := range text {
		if ch == 'a' {
			arr[0]++
		} else if ch == 'b' {
			arr[1]++
		} else if ch == 'n' {
			arr[2]++
		} else if ch == 'l' {
			arr[3]++
		} else if ch == 'o' {
			arr[4]++
		}
	}
	if arr[3] != 0 && arr[4] != 0 {
		arr[3] = arr[3] / 2
		arr[4] = arr[4] / 2
	}
	ans := arr[0]
	for _, e := range arr {
		if ans > e {
			ans = e
		}
	}
	return ans
}
