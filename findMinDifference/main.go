package main

import (
	"fmt"
	"math"
	"strings"
)

const MinutesOfDay = 24 * 60

func main() {
	timePoints := []string{"12:12", "00:13"}
	fmt.Println(findMinDifference(timePoints))
}

func findMinDifference(timePoints []string) int {
	bitSet := make([]bool, MinutesOfDay)
	for _, time := range timePoints {
		timeSplit := strings.Split(time, ":")
		h, m := timeSplit[0], timeSplit[1]
		totalMin := ((h[0]-'0')*10+(h[1]-'0'))*60 + (m[0]-'0')*10 + (m[1] - '0')
		if bitSet[totalMin] {
			return 0
		}
		bitSet[totalMin] = true
	}

	var prevMin, firstMin, lastMin, ans int = math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32

	for i := 0; i < MinutesOfDay; i++ {

		if !bitSet[i] {
			continue
		}
		if prevMin != math.MaxInt32 {
			if ans > i-prevMin {
				ans = i - prevMin
			}
		}
		prevMin = i
		if firstMin == math.MaxInt32 {
			firstMin = i
		}
		lastMin = i
	}

	if MinutesOfDay-lastMin+firstMin < ans {
		return MinutesOfDay - lastMin + firstMin
	}
	return ans
}
