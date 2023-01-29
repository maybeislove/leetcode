package main

import "strings"

func main() {
	//println(replaceSpaces("Mr John Smith    ", 13))
	println(replaceSpaces("               ", 5))

}

func replaceSpaces(S string, length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteString(string(S[i]))
		}
	}
	return sb.String()
}
