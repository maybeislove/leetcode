package main

func main() {
	//println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	//println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	println(hasGroupsSizeX([]int{1, 1}))

}

func hasGroupsSizeX(deck []int) bool {
	count := make([]int, 10000)
	for _, d := range deck {
		count[d]++
	}

	g := 0
	for i := 0; i < 10000; i++ {
		if count[i] > 0 {
			if g == -1 {
				g = count[i]
			} else {
				g = gcd(g, count[i])
			}
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
