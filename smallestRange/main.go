package main

import "math"

type MinElement struct {
	value   int
	ListIdx int
	ElIdx   int
}

type PriorityQueue []*MinElement

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq *PriorityQueue) Push(el any) {
	item := el.(*MinElement)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func smallestRange(nums [][]int) []int {
	pq := make(PriorityQueue, 0, len(nums))
	currMax := math.MinInt32
	for i, num := range nums {
		minValue := &MinElement{num[0], i, 0}
		currMax = max(currMax, num[0])
		pq.Push(minValue)
	}
	smallestRange := []int{0, math.MaxInt32}
	for pq.Len() == len(nums) {
		currMinEl := pq.Pop().(*MinElement)
		currMin, currListIdx, currElIdx := currMinEl.value, currMinEl.ListIdx, currMinEl.ElIdx
		diffCurrRange := currMax - currMin
		diffSmallestRange := smallestRange[1] - smallestRange[0]
		if diffCurrRange < diffSmallestRange {
			smallestRange[0] = currMin
			smallestRange[1] = currMax
		}

		if currElIdx+1 < len(nums[currListIdx]) {
			newVal := nums[currListIdx][currElIdx+1]
			newMinEl := &MinElement{newVal, currListIdx, currElIdx + 1}
			pq.Push(newMinEl)
			currMax = max(currMax, newVal)
		}
	}
	return smallestRange
}
