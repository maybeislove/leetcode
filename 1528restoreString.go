package main

func main() {
	println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}

func restoreString(s string, indices []int) string {
	b := make([]byte, len(s))

	for i, _ := range s {
		b[indices[i]] = s[i]
	}
	return string(b)
}
