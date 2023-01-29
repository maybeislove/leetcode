package main

func main() {
	println(titleToNumber("AA"))
	println(titleToNumber("AB"))
	println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {
	result := 0
	for i, _ := range columnTitle {
		temp := int(columnTitle[i] - 'A' + 1)
		result = result*26 + temp
	}
	return result
}
