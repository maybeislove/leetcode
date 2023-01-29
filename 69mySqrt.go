package main

func main() {
	println(mySqrt(4))
	mySqrt2(25)
}

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	if x < 3 {
		return 1
	}
	i := 2
	for i*i <= x {
		i++
	}
	return i - 1
}

func mySqrt2(x int) int {
	left, right := 0, x

	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
