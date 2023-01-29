package main

func main() {
	diStringMatch("DDI")
}
func diStringMatch(s string) []int {
	ints := make([]int, len(s)+1)
	max, min := len(s), 0
	for i, ch := range s {
		if ch == 'D' {
			ints[i] = max
			max--
		} else {
			ints[i] = min
			min++
		}
	}
	ints[len(s)] = max
	return ints
}
