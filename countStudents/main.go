package main

func countStudents(students []int, sandwiches []int) int {
	squareSt, circleSt := 0, 0
	for _, st := range students {
		if st == 0 {
			circleSt++
		} else {
			squareSt++
		}
	}

	for _, sw := range sandwiches {
		switch sw {
		case 0:
			if circleSt == 0 {
				return squareSt
			}
			circleSt--
		case 1:
			if squareSt == 0 {
				return circleSt
			}
			squareSt--
		}
	}
	return 0
}
