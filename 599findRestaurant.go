package main

import (
	"math"
)

func main() {
	findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"})
}

func findRestaurant(list1 []string, list2 []string) []string {

	result := []string{}
	m := map[string]int{}

	for i, s := range list1 {
		m[s] = i
	}
	indexSum := math.MaxInt32
	for i, s := range list2 {
		if j, ok := m[s]; ok {
			if i+j < indexSum {
				indexSum = i + j
				result = []string{s}
			} else if i+j == indexSum {
				result = append(result, s)
			}
		}

	}
	return result
}
