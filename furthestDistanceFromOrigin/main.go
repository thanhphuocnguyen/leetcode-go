package main

import "fmt"

func furthestDistanceFromOrigin(moves string) int {
	l, r, empty := 0, 0, 0
	for _, ru := range moves {
		switch ru {
		case 'L':
			l++
		case 'R':
			r++
		default:
			empty++
		}
	}

	return abs(l-r) + empty
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(furthestDistanceFromOrigin("L_RL__R"))
}
