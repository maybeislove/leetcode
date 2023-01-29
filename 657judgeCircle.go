package main

func main() {
	println(judgeCircle("DU"))
}

func judgeCircle(moves string) bool {

	vertical := 0
	horizontal := 0
	for _, ch := range moves {
		if ch == 'U' {
			vertical++
		} else if ch == 'D' {
			vertical--
		} else if ch == 'R' {
			horizontal++
		} else if ch == 'L' {
			horizontal--
		}
	}
	return vertical == 0 && horizontal == 0
}
