package main

func main() {
	println(rotateString("abcde", "cdeab"))
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if goal != s {
			s = s[1:] + string(s[0])
			continue
		}
		return true
	}
	return false
}
