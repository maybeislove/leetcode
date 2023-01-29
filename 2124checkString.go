package main

import "math"

func main() {
	println(checkString("aaabbb"))
	println(checkString("abab"))
	println(checkString("bbb"))
}

func checkString(s string) bool {
	a, b := 0, math.MaxInt32

	for i, ch := range s {
		if ch == 'a' && i >= a {
			a = i
		} else if i < b {
			b = i
		}
	}
	return b >= a
}
