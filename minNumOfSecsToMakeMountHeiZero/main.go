package main

import (
	"fmt"
	"math"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	// Better upper bound: worst case is one worker does everything
	minWorkerTime := int64(workerTimes[0])
	for _, t := range workerTimes {
		if int64(t) < minWorkerTime {
			minWorkerTime = int64(t)
		}
	}

	var lo, hi int64 = 1, minWorkerTime * int64(mountainHeight) * int64(mountainHeight+1) / 2

	for lo < hi {
		mid := lo + (hi-lo)/2
		if possible(workerTimes, int64(mountainHeight), mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func possible(workerTimes []int, mountainHeight, midTime int64) bool {
	// For each worker, calculate how much height they can reduce in midTime
	// If worker has baseTime t and needs to reduce height h:
	// total time = t * (1 + 2 + 3 + ... + h) = t * h * (h + 1) / 2
	// So: midTime >= t * h * (h + 1) / 2
	// Rearranging: h^2 + h - 2*midTime/t <= 0
	// Using quadratic formula: h <= (-1 + sqrt(1 + 8*midTime/t)) / 2

	var totalHeight int64 = 0
	for _, workerTime := range workerTimes {
		baseTime := int64(workerTime)

		// Calculate maximum height this worker can reduce
		// Using integer arithmetic to avoid floating point errors
		// h = floor((-1 + sqrt(1 + 8*midTime/baseTime)) / 2)

		if midTime < baseTime {
			// Worker can't even do 1 unit of work
			continue
		}

		// Calculate 8 * midTime / baseTime, but handle potential overflow
		eightTimesRatio := (midTime / baseTime) * 8
		if midTime%baseTime != 0 {
			eightTimesRatio += (midTime % baseTime) * 8 / baseTime
		}

		// Calculate sqrt(1 + 8*midTime/baseTime)
		discriminant := 1 + eightTimesRatio
		sqrtDiscriminant := int64(math.Sqrt(float64(discriminant)))

		// Make sure we don't underestimate due to floating point truncation
		if sqrtDiscriminant*sqrtDiscriminant < discriminant {
			sqrtDiscriminant++
		}
		if sqrtDiscriminant*sqrtDiscriminant > discriminant {
			sqrtDiscriminant--
		}

		// h = (sqrtDiscriminant - 1) / 2
		if sqrtDiscriminant > 1 {
			maxHeight := (sqrtDiscriminant - 1) / 2
			totalHeight += maxHeight

			if totalHeight >= mountainHeight {
				return true
			}
		}
	}

	return totalHeight >= mountainHeight
}

func main() {
	fmt.Println(minNumberOfSeconds(3, []int{2, 1, 1}))
}
