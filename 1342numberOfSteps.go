package main

func main() {
	println(numberOfSteps(14))
	println(numberOfSteps(8))
	println(numberOfSteps(123))

}

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 1 {
			count++
			if num > 1 {
				num -= 1
				count++
			}
		} else {
			count++
		}
		num /= 2
	}
	return count
}
