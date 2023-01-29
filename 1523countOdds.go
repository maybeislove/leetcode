package main

func main() {
	println(countOdds(3, 7))

}

func countOdds(low int, high int) int {

	return (high+1)/2 - low/2
}
