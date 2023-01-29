package main

func main() {
	largestAltitude([]int{-5, 1, 5, 0, -7})
}

func largestAltitude(gain []int) int {
	a := make([]int, len(gain)+1)
	a[0] = 0
	for i := 1; i <= len(gain); i++ {
		a[i] = a[i-1] + gain[i-1]
	}
	result := 0
	for _, e := range a {
		if e > result {
			result = e
		}
	}
	return result
}
