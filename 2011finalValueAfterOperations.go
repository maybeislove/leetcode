package main

import "strings"

func main() {

}

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, operation := range operations {
		if strings.Contains(operation, "+") {
			x++
		} else {
			x--
		}
	}
	return x
}
