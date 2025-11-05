package main

import (
	"container/heap"
	"math"
)

var moveDirections = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	visited := make([][]bool, n)
	dst := make([][]int, n)
	for i := range n {
		visited[i] = make([]bool, m)
		dst[i] = make([]int, m)
		for j := range m {
			dst[i][j] = math.MaxInt32
		}
	}
	dst[0][0] = 0
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		row:      0,
		col:      0,
		priority: 0,
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		currRow, currCol := item.row, item.col
		if visited[currRow][currCol] {
			continue
		}
		visited[currRow][currCol] = true
		for _, dir := range moveDirections {
			rowMove, colMove := dir[0], dir[1]
			nextRow, nextCol := currRow+rowMove, currCol+colMove
			if !checkBound(nextRow, nextCol, n, m) {
				continue
			}
			moveTime := max(moveTime[nextRow][nextCol], dst[currRow][currCol]) + 1
			if moveTime < dst[nextRow][nextCol] {
				dst[nextRow][nextCol] = moveTime
				heap.Push(&pq, &Item{
					row:      nextRow,
					col:      nextCol,
					priority: moveTime,
				})
			}
		}
	}
	return dst[n-1][m-1]
}

func checkBound(i, j, n, m int) bool {
	return i >= 0 && j >= 0 && i < n && j < m
}

// An Item is something we manage in a priority queue.
type Item struct {
	row      int // The value of the item; arbitrary.
	col      int // The index of the item in the heap.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
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
	// Example usage
	moveTime := [][]int{
		{0, 4},
		{4, 4},
	}
	result := minTimeToReach(moveTime)
	println(result) // Output: 6
}
