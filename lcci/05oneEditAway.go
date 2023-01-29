package main

import "math"

func main() {
	println(oneEditAway("aabb", "abba"))
}

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if math.Abs(float64(m)-float64(n)) > 1 {
		return false
	}
	if n > m {
		return oneEditAway(second, first)
	}
	for i, _ := range second {
		if first[i] != second[i] {
			if m == n {
				return first[i+1:] == second[i+1:]
			} else {
				return first[i+1:] == second[i:]
			}
		}
	}
	return true
}
