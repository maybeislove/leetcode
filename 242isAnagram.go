package main

import (
	"sort"
	"strings"
)

func main() {
	println(isAnagram("faa", "aaf"))
}

func isAnagram(s string, t string) bool {
	s = SortString(s)
	t = SortString(t)
	return s == t
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isAnagram2(s string, t string) bool {

	c1 := [26]int{}
	c2 := [26]int{}
	for _, v := range s {
		c1[v-'a']++
	}
	for _, v := range t {
		c2[v-'a']++
	}
	return c1 == c2
}
