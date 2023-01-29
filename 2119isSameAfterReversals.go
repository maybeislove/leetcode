package main

func main() {
	println(isSameAfterReversals(526))
}

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	return num%10 != 0
}
