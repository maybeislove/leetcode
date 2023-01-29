package main

import "fmt"

func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}

func repeatedCharacter(s string) byte {
	m := make(map[byte]struct{}, len(s))
	for _, ch := range s {
		if _, ok := m[byte(ch)]; ok {
			return byte(ch)
		}
		m[byte(ch)] = struct{}{}
	}
	return 0
}
