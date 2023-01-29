package main

import (
	"fmt"
)

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
}

func garbageCollection(garbage []string, travel []int) int {
	g, p, m := 0, 0, 0
	gc, pc, mc := 0, 0, 0
	for i, s := range garbage {
		for _, ch := range s {
			if ch == 'G' {
				g = i
				gc++
			} else if ch == 'P' {
				pc++
				p = i
			} else {
				mc++
				m = i
			}
		}
	}
	result := gc + pc + mc
	for i, val := range travel {
		if g > i {
			result += val
		}
		if p > i {
			result += val
		}
		if m > i {
			result += val
		}
	}
	return result
}

func max4(a, b int) int {
	if a > b {
		return a
	}
	return b
}
