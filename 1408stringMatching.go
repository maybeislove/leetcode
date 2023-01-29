package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("%v", stringMatching([]string{"mass", "as", "hero", "superhero"}))
}

func stringMatching(words []string) []string {
	var m = make(map[string]bool, len(words))

	if len(words) < 1 {
		return []string{}
	}
	result := []string{}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	for i, _ := range words {
		for j := i + 1; j < len(words); j++ {
			if len(words[i]) < len(words[j]) {
				continue
			}
			if strings.Contains(words[i], words[j]) {
				if m[words[j]] {
					continue
				}
				m[words[j]] = true
				result = append(result, words[j])
			}
		}
	}
	return result
}
