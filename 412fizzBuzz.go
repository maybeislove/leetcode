package main

import "fmt"

func main() {
	fmt.Printf("%v", fizzBuzz(15))
}
func fizzBuzz(n int) []string {
	res := []string{}
	var fizz, buzz, fizzbuzz = "Fizz", "Buzz", "FizzBuzz"
	for i := 1; i <= n; i++ {
		word := ""
		if i%3 == 0 && i%5 == 0 {
			word = fizzbuzz
		} else if i%3 == 0 {
			word = fizz
		} else if i%5 == 0 {
			word = buzz
		} else {
			word = fmt.Sprintf("%d", i)
		}
		res = append(res, word)
	}
	return res
}
