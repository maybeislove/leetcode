package main

func main() {
	println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	m := map[int32]int{}
	for _, a := range magazine {
		m[a-'a']++
	}
	for _, a := range ransomNote {
		m[a-'a']--
		if m[a-'a'] < 0 {
			return false
		}
	}
	return true
}
