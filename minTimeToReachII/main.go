package main

import (
	"container/heap"
	"math"
)

var dirs = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	dst := make([][]int, n)
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
		dst[i] = make([]int, m)
		for j := range m {
			dst[i][j] = math.MaxInt32
		}
	}
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		moveCost: 1,
		row:      0,
		col:      0,
		priority: 0,
	}
	heap.Init(&pq)
	dst[0][0] = 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		currRow, currCol, moveCost := item.row, item.col, item.moveCost
		if visited[currRow][currCol] {
			continue
		}
		visited[currRow][currCol] = true
		for _, dir := range dirs {
			rowMove, colMove := dir[0], dir[1]
			nextRow, nextCol := currRow+rowMove, currCol+colMove
			if nextRow < 0 || nextCol < 0 || nextRow >= n || nextCol >= m {
				continue
			}
			nextTimeCost := max(moveTime[nextRow][nextCol], dst[currRow][currCol]) + moveCost
			if nextTimeCost < dst[nextRow][nextCol] {
				nextMoveCost := moveCost
				if moveCost == 1 {
					nextMoveCost = 2
				} else {
					nextMoveCost = 1
				}
				dst[nextRow][nextCol] = nextTimeCost
				heap.Push(&pq, &Item{
					moveCost: nextMoveCost,
					row:      nextRow,
					col:      nextCol,
					priority: nextTimeCost,
				})
			}
		}
	}
	return dst[n-1][m-1]
}

// An Item is something we manage in a priority queue.
type Item struct {
	moveCost int
	row      int
	col      int
	priority int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func main() {
	println(minTimeToReach([][]int{{0, 58}, {27, 69}}))
}
