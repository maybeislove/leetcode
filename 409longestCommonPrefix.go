package main

import (
	"strings"
)

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
}

func longestCommonPrefix(strs []string) string {
	s := strs[0]
	if len(s) == 0 {
		return ""
	}
	for _, str := range strs {
		for !strings.HasPrefix(str, s) {
			if len(s) == 0 {
				return ""
			}
			s = s[0 : len(s)-1]
			println(s)
		}
	}
	return s
}
