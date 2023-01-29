package main

import "fmt"

func main() {
	//println(convertToTitle(1))
	fmt.Printf(convertToTitle(28))

}
func convertToTitle(columnNumber int) string {
	var result []byte

	for columnNumber > 0 {
		columnNumber -= 1
		result = append(result, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
