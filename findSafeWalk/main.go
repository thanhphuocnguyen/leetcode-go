package main

import (
	"container/heap"
	"fmt"
	"math"
)

var DIRECTIONS = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	// declare dist
	healthNeeded := make([][]int, m)
	for i := range m {
		healthNeeded[i] = make([]int, n)
		for j := range n {
			healthNeeded[i][j] = math.MaxInt32
		}
	}
	healthNeeded[0][0] = grid[0][0]
	// use bfs to traverse find the shortest path
	pq := PriorityQueue{
		&Item{
			r:        0,
			c:        0,
			consumed: grid[0][0],
		},
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)
		r, c, consumed := curr.r, curr.c, curr.consumed
		if r == m-1 && c == n-1 {
			break
		}
		if consumed > healthNeeded[r][c] {
			continue
		}
		for _, dir := range DIRECTIONS {
			nr, nc := r+dir[0], c+dir[1]
			if isValid(nr, nc, m, n) && consumed+grid[nr][nc] < healthNeeded[nr][nc] {
				healthNeeded[nr][nc] = consumed + grid[nr][nc]
				heap.Push(&pq, &Item{
					r:        nr,
					c:        nc,
					consumed: healthNeeded[nr][nc],
				})
			}
		}
	}
	// get the shortest path - health and return
	return health-healthNeeded[m-1][n-1] > 0
}

func isValid(r, c, m, n int) bool {
	return r >= 0 && c >= 0 && r < m && c < n
}

// An Item is something we manage in a priority queue.
type Item struct {
	r        int
	c        int
	consumed int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].consumed < pq[j].consumed
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
	fmt.Println(findSafeWalk([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}, 1))
}
