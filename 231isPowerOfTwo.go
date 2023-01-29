package main

import (
	"strconv"
	"strings"
)

func main() {
	println(isPowerOfTwo(-16))
}
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	formatInt := strconv.FormatInt(int64(n), 2)
	println(formatInt)

	split := strings.Split(formatInt, "1")
	return len(split) == 2
}
func isPowerOfTwo2(n int) bool {
	return n > 0 && n&(n-1) == 0
}
