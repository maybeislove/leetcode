package main

func main() {

}

func maximumWealth(accounts [][]int) int {
	result := 0
	for _, account := range accounts {
		sum := 0
		for _, a := range account {
			sum += a
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
