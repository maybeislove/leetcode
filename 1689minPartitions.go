package main

import "fmt"

func main() {
	fmt.Println(minPartitions("27346209830709182346"))
}
func minPartitions(n string) int {
	mx := 0
	for _, ch := range n {
		mx = max3(int(ch-'0'), mx)
	}
	return mx
}
func max3(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
