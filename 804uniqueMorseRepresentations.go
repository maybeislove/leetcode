package main

func main() {
	println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

func uniqueMorseRepresentations(words []string) int {
	var morse = []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-",
		"...-", ".--", "-..-", "-.--", "--..",
	}
	m := make(map[string]struct{})

	for _, word := range words {
		temp := ""
		for _, ch := range word {
			temp += morse[ch-'a']
		}
		m[temp] = struct{}{}
	}
	return len(m)
}
