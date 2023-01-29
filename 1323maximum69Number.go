package main

import "strconv"

func main() {
	println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	s := strconv.Itoa(num)
	b := []byte(s)
	for i, ch := range b {
		if ch == '6' {
			b[i] = '9'
			break
		}
	}
	a, _ := strconv.Atoi(string(b))
	return a
}
