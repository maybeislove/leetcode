package main

func main() {
	println(canPlaceFlowers([]int{0, 1, 0, 0, 1, 0, 0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
		} else if len(flowerbed)-1 == i || flowerbed[i+1] == 0 {
			n--
			i++
		} else if flowerbed[i+1] == 1 {
			i += 2
		}
		if n == 0 {
			return true
		}
	}
	return n <= 0
}
