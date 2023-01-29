package main

func main() {
	println(numJewelsInStones("aA", "aAAbbbb"))
}
func numJewelsInStones(jewels string, stones string) int {

	m := make(map[int32]struct{}, len(jewels))

	for _, ch := range jewels {
		m[ch] = struct{}{}
	}
	result := 0
	for _, ch := range stones {
		if _, ok := m[ch]; ok {
			result++
		}
	}
	return result
}
