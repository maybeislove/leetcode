package main

func main() {
	println(balancedStringSplit("RLRRRLLRLL"))
}

func balancedStringSplit(s string) int {

	l, r, res := 0, 0, 0

	for _, char := range s {
		if char == 'L' {
			l++
		} else {
			r++
		}
		if l == r {
			res++
		}
	}
	return res
}
