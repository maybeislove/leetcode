package main

func main() {

}
func guessNumber(n int) int {

	var l, r = 0, n
	for l < r {
		mid := (l + r) / 2
		i := guess(mid)
		if i > 0 {
			l = mid + 1
		} else if i < 0 {
			r = mid - 1
		} else {
			return i
		}
	}
	return l
}

func guess(num int) int {
	return 0
}
