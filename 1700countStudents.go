package main

func main() {
	students := []int{1, 1, 0, 0, 1}
	sandwiches := []int{0, 1, 0, 1, 0}
	println(countStudents(students, sandwiches))
}

func countStudents(students []int, sandwiches []int) int {
	s0 := 0
	s1 := 0
	for _, student := range students {
		if student == 0 {
			s0++
		} else {
			s1++
		}
	}
	for _, sandwich := range sandwiches {

		if sandwich == 0 && s0 > 0 {
			s0--
		} else if sandwich == 1 && s1 > 0 {
			s1--
		} else {
			break
		}
	}

	return s0 + s1
}
