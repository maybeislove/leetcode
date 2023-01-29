package main

func main() {
	println(isSubsequence("abc", "ahbgdc"))
}
func isSubsequence(s string, t string) bool {

	var l, r, m, n = 0, 0, len(s), len(t)
	for l < m && r < n {
		if s[l] == t[r] {
			l++
		}
		r++
	}
	return l == m
}
