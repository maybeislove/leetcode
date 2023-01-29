package main

func main() {

}

func minSwaps(s string) int {
	result := 0
	temp := 0
	for _, ch := range s {
		if ch == ']' && temp == 0 {
			result++
			temp++
		} else if ch == '[' {
			temp++
		} else if ch == ']' {
			temp--
		}
	}
	return result
}
