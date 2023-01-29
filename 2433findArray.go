package main

import "fmt"

func main() {
	fmt.Printf("%v", findArray([]int{5, 2, 0, 3, 1}))
}
func findArray(pref []int) []int {
	result := []int{pref[0]}
	for i := 1; i < len(pref); i++ {
		result = append(result, pref[i]^pref[i-1])
	}
	return result
}
