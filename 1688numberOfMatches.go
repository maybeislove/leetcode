package main

func main() {
	println(numberOfMatches(14))
}

func numberOfMatches(n int) int {
	result := 0
	for n != 1 {
		if n%2 == 0 {
			result += n / 2
			n /= 2
		} else {
			result += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return result
}
