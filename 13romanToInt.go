package main

func main() {
	println(romanToInt("IX"))
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0
	for i, ch := range s {
		if i < len(s)-1 && roman[string(s[i])] < roman[string(s[i+1])] {
			result -= roman[string(ch)]
		} else {
			result += roman[string(ch)]
		}
	}
	return result
}
