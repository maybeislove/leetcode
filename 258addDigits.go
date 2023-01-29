package main

func main() {
	println(addDigits(38))
}
func addDigits(num int) int {
	for num >= 10 {
		a := 0
		for num > 0 {
			a += num % 10
			num /= 10
		}
		num = a
	}
	return num
}
