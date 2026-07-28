package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

var DIRECTIONS = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}

	safenessScores := computeSafenessScores(n, grid)

	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = -1
		}
	}
	dist[0][0] = safenessScores[0][0]
	pq := PriorityQueue{
		&Item{
			r:     0,
			c:     0,
			score: dist[0][0],
		},
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*Item)
		r, c, score := cell.r, cell.c, cell.score
		if r == n-1 && c == n-1 {
			return score
		}
		if score < dist[r][c] {
			continue
		}
		for _, dir := range DIRECTIONS {
			nr, nc := r+dir[0], c+dir[1]
			if isValid(nr, nc, n) {
				newScore := min(score, safenessScores[nr][nc])
				if newScore > dist[nr][nc] {
					dist[nr][nc] = newScore
					heap.Push(&pq, &Item{
						r:     nr,
						c:     nc,
						score: newScore,
					})
				}
			}
		}
	}

	return 0
}

func isValid(r, c, n int) bool {
	return r >= 0 && c >= 0 && r < n && c < n
}

func computeSafenessScores(n int, grid [][]int) [][]int {
	safeness := make([][]int, n)
	queue := list.New()
	for i := range n {
		safeness[i] = make([]int, n)
		for j := range n {
			if grid[i][j] == 1 {
				safeness[i][j] = 0
				queue.PushBack([2]int{i, j})
			} else {
				safeness[i][j] = math.MaxInt32
			}
		}
	}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		thief := front.Value.([2]int)
		r, c := thief[0], thief[1]
		for _, dir := range DIRECTIONS {
			nr, nc := r+dir[0], c+dir[1]
			if isValid(nr, nc, n) && safeness[nr][nc] == math.MaxInt32 {
				safeness[nr][nc] = safeness[r][c] + 1
				queue.PushBack([2]int{nr, nc})
			}
		}
	}

	return safeness
}

// An Item is something we manage in a priority queue.
type Item struct {
	r     int // The value of the item; arbitrary.
	c     int // The value of the item; arbitrary.
	score int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].score > pq[j].score
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
	fmt.Println(maximumSafenessFactor([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))
}
