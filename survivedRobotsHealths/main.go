package main

import (
	"fmt"
	"sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	// define robot at index
	robotIdx := make([]int, len(positions))
	for i := range len(positions) {
		robotIdx[i] = i
	}

	// sort robot's indexes by their positions
	sort.Slice(robotIdx, func(i, j int) bool {
		return positions[robotIdx[i]] < positions[robotIdx[j]]
	})
	// track survived robots go to right
	surviveRobot := make([]int, 0)
	ans := make([]int, 0)
	for _, currRobot := range robotIdx {
		// if current is right so add it to survived robots
		if directions[currRobot] == 'R' {
			surviveRobot = append(surviveRobot, currRobot)
		} else {
			// if current robot is going to the left we compare with the right most survived robot
			for len(surviveRobot) > 0 && healths[currRobot] > 0 {
				rightIdx := surviveRobot[len(surviveRobot)-1]
				if healths[rightIdx] == healths[currRobot] {
					// remove both robots out of the line
					surviveRobot = surviveRobot[:len(surviveRobot)-1]
					healths[rightIdx] = 0
					healths[currRobot] = 0
				} else if healths[rightIdx] > healths[currRobot] {
					// if go right robot > go left remove currRobot out of the line, keep the go right but decrease health 1
					healths[rightIdx]--
					healths[currRobot] = 0
				} else {
					// if go left > go right robot remove survived robot out of stack and continue compare the surviving robots, decrease health go left 1
					healths[currRobot]--
					healths[rightIdx] = 0
					surviveRobot = surviveRobot[:len(surviveRobot)-1]
				}
			}
		}
	}

	for _, h := range healths {
		if h > 0 {
			ans = append(ans, h)
		}
	}
	return ans
}

func main() {
	fmt.Println(survivedRobotsHealths([]int{3, 5, 2, 6}, []int{10, 10, 15, 12}, "RLRL"))
	fmt.Println(survivedRobotsHealths([]int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}, "RRRRR"))
}
