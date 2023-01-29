package main

import "fmt"

func main() {
	fmt.Printf("%v", largeGroupPositions("abbxxxxzyy"))
	fmt.Printf("%v", largeGroupPositions("abcdddeeeeaabbbcd"))
}

func largeGroupPositions(s string) [][]int {

	result := [][]int{}
	count := 1
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if count >= 3 {
				result = append(result, []int{i - count + 1, i})
			}
			count = 1
		} else {
			count++
		}
	}
	return result
}
