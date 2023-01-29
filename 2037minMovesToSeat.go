package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	result := 0
	for i, seat := range seats {
		result += int(math.Abs(float64(students[i] - seat)))
	}
	return result
}
