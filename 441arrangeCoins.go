package main

import "math"

func main() {
	println(arrangeCoins(4))
}

//前n項和公式為：Sn=na1+n(n-1)d/2
//或Sn=n(a1+an)/2
//注意：
//以上n均屬於正整數。
func arrangeCoins(n int) int {
	return int((-1 + math.Sqrt(float64(1+(8*n)))) / 2)
}
