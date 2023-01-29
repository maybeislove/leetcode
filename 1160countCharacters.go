package main

func main() {
	println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {
	m := make(map[int32]int, len(chars))
	for _, ch := range chars {
		m[ch]++
	}
	result := 0
	for _, word := range words {
		n := make(map[int32]int, len(word))
		f := true
		for _, ch := range word {
			n[ch]++
		}
		for k, v := range n {
			if v > m[k] {
				f = false
				break
			}
		}
		if f {
			result += len(word)
		}
	}
	return result
}
