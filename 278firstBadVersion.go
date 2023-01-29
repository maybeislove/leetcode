package main

import "math/rand"

func main() {

}
func firstBadVersion(n int) int {
	var left, right = 0, n
	for left < right {
		mid := left + (right-left)/2
		println(mid)
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return rand.Int()%3 == 0
}
