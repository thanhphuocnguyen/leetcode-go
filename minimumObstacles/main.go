package main

import (
	"container/heap"
	"math"
)

type Item [3]int
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i][0] > pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minimumObstacles(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	directions := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = grid[0][0]
	pq := make(PriorityQueue, 1)
	heap.Init(&pq)
	heap.Push(&pq, Item{
		0, 0, 0,
	})
	for pq.Len() > 0 {
		currItem := heap.Pop(&pq).(Item)
		weight, row, col := currItem[0], currItem[1], currItem[2]
		if row == n-1 && col == m-1 {
			return weight
		}
		for _, dir := range directions {
			nextRow, nextCol := row+dir[0], col+dir[1]
			if isValid(n, m, nextRow, nextCol) {
				nextWeight := weight + grid[nextRow][nextCol]
				if nextWeight < dist[nextRow][nextCol] {
					dist[nextRow][nextCol] = nextWeight
					heap.Push(&pq, Item{nextWeight, nextRow, nextCol})
				}
			}
		}
	}
	return dist[n-1][m-1]
}

func isValid(n, m, i, j int) bool {
	if i >= 0 && i < n && j >= 0 && j < m {
		return true
	}
	return false
}

func main() {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	minimumObstacles(grid)
}
