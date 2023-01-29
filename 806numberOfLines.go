package main

import "fmt"

func main() {
	fmt.Printf("%v", numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
}

func numberOfLines(widths []int, s string) []int {
	resultA := 1
	rowWidth := 100
	for _, ch := range s {

		if rowWidth-widths[ch-'a'] >= 0 {
			rowWidth -= widths[ch-'a']
		} else {
			resultA++
			rowWidth = 100 - widths[ch-'a']
		}
	}
	return []int{resultA, 100 - rowWidth}
}
