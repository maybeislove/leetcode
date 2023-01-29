package main

func main() {
	println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}

func validPalindrome(s string) bool {
	var l, r = 0, len(s) - 1

	for l < r {

		if s[l] == s[r] {
			l++
			r--
			continue
		} else {
			flag1, flag2 := true, true
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}

			}
			for i, j := l+1, r; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
