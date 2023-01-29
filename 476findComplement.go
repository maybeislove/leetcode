package main

func main() {

}

func findComplement(num int) int {
	highbit := 0
	for i := 1; i <= 31; i++ {

		if num < 1<<i {
			break
		}
		highbit = i
	}
	mask := 1<<(highbit+1) - 1
	return num ^ mask
}
