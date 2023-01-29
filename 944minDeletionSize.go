package main

func main() {
	println(minDeletionSize([]string{"abc", "bce", "cae"}))
}

func minDeletionSize(strs []string) int {
	ans := 0
	for j := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] > strs[i][j] {
				ans++
				break
			}
		}
	}
	return ans
}
