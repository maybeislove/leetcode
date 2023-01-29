package main

import (
	"fmt"
	"strings"
)

func main() {
	println(oddString([]string{"adc", "wzy", "abc"}))
}

func oddString(words []string) string {
	m := make(map[string][]string, 0)
	for _, word := range words {
		val := []string{}
		for i := 1; i < len(word); i++ {
			val = append(val, fmt.Sprintf("%v", word[i]-word[i-1]))
		}
		key := strings.Join(val, ",")
		m[key] = append(m[key], word)
	}
	for _, v := range m {
		if len(v) == 1 {
			return v[0]
		}
	}
	return ""
}
