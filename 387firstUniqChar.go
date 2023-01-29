package main

func main() {
	println(firstUniqChar("abcvcab"))
}
func firstUniqChar(s string) int {

	m := make(map[int32]int, 26)

	for _, ch := range s {
		m[ch-'a']++
	}

	for i, ch := range s {
		if m[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
