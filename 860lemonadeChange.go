package main

func main() {
	println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	println(lemonadeChange([]int{5, 5, 5, 20}))
}

func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}
	if bills[0] != 5 {
		return false
	}

	doller5, doller10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			doller5++
		} else if bill == 10 {
			if doller5 < 1 {
				return false
			}
			doller5--
			doller10++
		} else {
			if doller5 > 0 && doller10 > 0 {
				doller5--
				doller10--
			} else if doller5 >= 3 {
				doller5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}
