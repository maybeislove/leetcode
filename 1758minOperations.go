package main

import "math"

func main() {
	println(minOperations("01110"))
}

func minOperations(s string) int {
	count := 0
	for i, ch := range s {
		if int(ch-'0') != i%2 {
			count++
		}
	}
	return int(math.Min(float64(count), float64(len(s)-count)))
}
