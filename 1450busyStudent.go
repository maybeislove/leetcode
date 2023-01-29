package main

func main() {
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	result := 0
	for i, _ := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			result++
		}
	}
	return result
}
