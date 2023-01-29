package main

func main() {

}
func isUnique(astr string) bool {

	m := make(map[int32]int, len(astr))
	for _, ch := range astr {
		m[ch]++
		if m[ch] > 1 {
			return false
		}
	}
	return true
}
