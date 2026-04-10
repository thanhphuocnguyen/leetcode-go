package main

import (
	"fmt"
	"sort"
)

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	robotDist := make(map[int]int)
	// save distances fro look up
	for i, rb := range robots {
		robotDist[rb] = distance[i]
	}
	sort.Ints(robots)
	sort.Ints(walls)
	dpLeft := make([]int, n)
	dpRight := make([]int, n)
	wallsBetweenRobots := make([]int, n)

	// loop through each robot it can fire to the left or right
	for i, rb := range robots {
		// process for the left of current robot
		// boundaries is from max(robot on the left of current robot, farthest distance bullet of robot) - to wall nearest robots
		// [bullet stop .. wall next to robot] ->robot here
		ubLeft := upperBound(walls, rb)
		var lbLeft int
		if i == 0 {
			// lW
			lbLeft = lowerBound(walls, rb-robotDist[rb])
		} else {
			lbTarget := max(robots[i-1]+1, rb-robotDist[rb])
			// lW
			lbLeft = lowerBound(walls, lbTarget)
		}
		dpLeft[i] = ubLeft - lbLeft

		//robot here <- [wall next to robot .. bullet stop]
		lbRight := lowerBound(walls, rb)
		var ubRight int
		if i == n-1 {
			ubRight = upperBound(walls, rb+robotDist[rb])
		} else {
			ubTarget := min(robots[i+1]-1, rb+robotDist[rb])
			lbRight = upperBound(walls, ubTarget)
		}

		dpRight[i] = ubRight - lbRight
		if i != 0 {
			// calculate walls between current robot with prev robot
			// prev robot [ .. ] curr robot
			lbPrevLeft := lowerBound(walls, robots[i-1])
			wallsBetweenRobots[i] = ubLeft - lbPrevLeft
		}
	}

	subLeft, subRight := dpLeft[0], dpRight[0]
	for i := 1; i < n; i++ {
		currLeft := max(subLeft+dpLeft[i], subRight-dpRight[i-1]+min(dpLeft[i]+dpRight[i-1], wallsBetweenRobots[i]))
		currRight := max(subLeft+dpRight[i], subRight+dpRight[i])
		subLeft = currLeft
		subRight = currRight
	}
	return max(subLeft, subRight)
}

func upperBound(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	robots := []int{4}
	distance := []int{3}
	walls := []int{1, 10}
	fmt.Println(maxWalls(robots, distance, walls))
}
