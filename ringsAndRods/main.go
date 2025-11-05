package main

func countPoints(rings string) int {
	counts := make([]int, 10)
	for i := 0; i+1 < len(rings); i += 2 {
		rod := rings[i+1] - '0'
		switch rings[i] {
		case 'R':
			counts[rod] |= 1
		case 'G':
			counts[rod] |= 2
		case 'B':
			counts[rod] |= 4
		}
	}
	ans := 0
	for _, v := range counts {
		if v == 7 {
			ans++
		}
	}
	return ans
}

func main() {
	countPoints("B0B6G0R6R0R6G9")
}
