package main

import (
	"strconv"
)

func main() {
	addBinary("10", "101")
}

func addBinary(a string, b string) string {
	carry := 0
	maxLength := len(a)
	result := ""
	if len(a) < len(b) {
		maxLength = len(b)
	}
	for i := 0; i < maxLength; i++ {
		if i < len(a) {
			carry += int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			carry += int(b[len(b)-i-1] - '0')
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}
