package main

import "fmt"

func main() {
	fmt.Printf(toHex(-1))
}

func toHex(num int) string {
	return fmt.Sprintf("%x", num)
}
