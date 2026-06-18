package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	startIdx := 0
	potential := 0
	for i, c := range cost {
		if gas[i]-c > potential {
			startIdx = i
			potential = gas[i] - c
		}
	}
	backIdx := startIdx + n
	maxGas := 0
	for startIdx < backIdx {
		idx := startIdx % n
		// fill gas
		maxGas += gas[idx]
		gas[idx] = 0
		// move to new gas
		maxGas -= cost[idx]
		if maxGas < 0 {
			return -1
		}
		startIdx++
	}

	return backIdx - n
}

func main() {
	fmt.Println(canCompleteCircuit([]int{3, 3, 4}, []int{3, 4, 4}))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
