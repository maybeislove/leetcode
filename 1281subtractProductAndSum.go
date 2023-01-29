package main

func main() {
	println(subtractProductAndSum(234))
	println(subtractProductAndSum(4421))

}

func subtractProductAndSum(n int) int {
	if n == 0 {
		return 0
	}
	product, sum := 1, 0
	for n > 0 {
		i := n % 10
		sum += i
		product *= i
		n /= 10
	}

	return product - sum
}
