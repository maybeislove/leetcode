package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", constructRectangle(37))
}

func constructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))

	for area%sqrt != 0 {
		sqrt--
	}

	return []int{area / sqrt, sqrt}
}
