package main

import "strings"

func main() {

}
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	builder1, builder2 := strings.Builder{}, strings.Builder{}
	for _, s := range word1 {
		builder1.WriteString(s)
	}
	for _, s := range word2 {
		builder2.WriteString(s)
	}
	return builder1.String() == builder2.String()
}
