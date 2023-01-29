package main

import "fmt"

func main() {
	//fmt.Println(countDigits(121))
	fmt.Println(countDigits(54))

}

func countDigits(num int) int {
	result := 0
	temp := num
	for temp > 0 {
		i := temp % 10
		if num%i == 0 {
			result++
		}
		temp /= 10
	}
	return result
}
