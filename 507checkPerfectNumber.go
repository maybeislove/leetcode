package main

func main() {
	println(checkPerfectNumber(28))
	println(checkPerfectNumber(7))
	println(checkPerfectNumber(3))

}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sumFactor := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sumFactor += d
			if d*d < num {
				sumFactor += num / d
			}
		}
	}
	return num == sumFactor
}
