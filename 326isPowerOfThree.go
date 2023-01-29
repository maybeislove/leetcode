package main

func main() {
	println(isPowerOfThree(4))
}
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}
