package main

func main() {
	findTheDifference("a", "ab")
}

func findTheDifference(s string, t string) byte {
	u := uint8(0)
	for i := range t {
		u += t[i]
	}
	for i := range s {
		u -= s[i]
	}
	return u
}
