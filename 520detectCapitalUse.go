package main

func main() {
	println(detectCapitalUse("USA"))
	println(detectCapitalUse("FlaG"))
	println(detectCapitalUse("mL"))

}
func detectCapitalUse(word string) bool {

	var lastLowerIndex, lastUpperIndex = -1, -1

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			lastLowerIndex = i
		} else {
			lastUpperIndex = i
		}
	}

	if lastLowerIndex >= 0 && lastUpperIndex >= 1 {
		return false
	}
	return true
}

//如果一个单词里存在小写字母，且大写字母的位置不为0，那么大写用法不正确
