package main

func main() {
	println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}
