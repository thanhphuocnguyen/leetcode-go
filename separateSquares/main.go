package main

import "fmt"

func main() {
	fmt.Println(separateSquares([][]int{{0, 0, 1}, {2, 2, 1}}))
}
func separateSquares(squares [][]int) float64 {
	// find the upper bound and total area of all squares
	maxY, totalArea := 0.0, 0.0
	for _, sq := range squares {
		bY, l := sq[1], sq[2]
		tY := float64(bY + l)
		if tY > maxY {
			maxY = tY
		}
		totalArea += float64(l * l)
	}

	check := func(midY float64) bool {
		area := 0.0
		// find area of all squares/cut-squares below midY
		for _, sq := range squares {
			y, l := float64(sq[1]), float64(sq[2])
			if y < midY {
				// if upper-l (height) of current square > l. take the height from y to midY, else get l
				overlap := min(midY-y, l)
				// calculate area
				area += overlap * l
			}
		}
		// if 2*area >= totalArea mean we can decrease midY, else we have to increase midY
		return 2*area >= totalArea
	}

	lo, hi := 0.0, maxY
	// Epsilon is 0,00005
	esp := 1e-5
	// run till diff of hi and lo < epsilon
	for abs(hi-lo) >= esp {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
