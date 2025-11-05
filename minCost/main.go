package main

import (
	"container/heap"
	"math"
)

type Item struct {
	value    [3]int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
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

func minCost(grid [][]int) int {
	numRows, numCols := len(grid), len(grid[0])
	dirs := [4][2]int{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	}
	minCost := make([][]int, numRows)
	for i := range minCost {
		minCost[i] = make([]int, numCols)
		for j := range numCols {
			minCost[i][j] = math.MaxInt32
		}
	}

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{[3]int{0, 0, 0}, 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curPos := heap.Pop(&pq).(*Item).value
		cost, curRow, curCol := curPos[0], curPos[1], curPos[2]
		if cost != minCost[curRow][curCol] {
			continue
		}
		for i, dir := range dirs {
			nextRow, nextCol := curRow+dir[0], curCol+dir[1]
			if isValid(numRows, numCols, nextRow, nextCol) {
				nextCost := minCost[curRow][curCol]
				if grid[curRow][curCol] != i+1 {
					nextCost++
				}
				if minCost[nextRow][nextCol] > nextCost {
					minCost[nextRow][nextCol] = nextCost
					heap.Push(&pq, &Item{[3]int{nextCost, nextRow, nextCol}, nextCost})
				}
			}
		}
	}
	return minCost[numRows-1][numCols-1]
}

func isValid(numRows, numCols, curRow, curCol int) bool {
	return curRow >= 0 && curCol >= 0 && curRow < numRows && curCol < numCols
}
