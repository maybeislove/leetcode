package main

func main() {
	println(checkRecord("PPALLL"))
}

func checkRecord(s string) bool {
	var l, a = 0, 0
	for _, ch := range s {
		if l >= 3 {
			return false
		}
		if ch == 'L' {
			l++
			continue
		} else if ch == 'A' {
			a++
		}
		l = 0
	}

	if l >= 3 || a >= 2 {
		return false
	}
	return true
}
