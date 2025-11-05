package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var rowDir = []int{-1, 0, 1, 0}
var colDir = []int{0, -1, 0, 1}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	qWithIdx := make([][2]int, len(queries))
	for i, q := range queries {
		qWithIdx[i] = [2]int{q, i}
	}

	sort.Slice(qWithIdx, func(i, j int) bool {
		return qWithIdx[i][0] < qWithIdx[j][0]
	})
	point := 0
	visited[0][0] = true
	ans := make([]int, len(queries))
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{
		Row:      0,
		Col:      0,
		Priority: grid[0][0],
	})

	for _, qi := range qWithIdx {
		q, index := qi[0], qi[1]
		for pq.Len() > 0 && pq.Peek().(*Item).Priority < q {
			item := heap.Pop(&pq).(*Item)
			point++
			for i := range 4 {
				nextRow, nextCol := item.Row+rowDir[i], item.Col+colDir[i]
				if isValidCell(nextRow, nextCol, visited) {
					visited[nextRow][nextCol] = true
					heap.Push(&pq, &Item{
						Row:      nextRow,
						Col:      nextCol,
						Priority: grid[nextRow][nextCol],
					})
				}
			}
		}

		ans[index] = point
	}
	return ans
}

func isValidCell(row, col int, visited [][]bool) bool {
	m, n := len(visited), len(visited[0])
	return row >= 0 && col >= 0 && row < m && col < n && !visited[row][col]
}

type Item struct {
	Row      int // The value of the item; arbitrary.
	Col      int // The index of the item in the heap.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
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
func (pq PriorityQueue) Peek() any { return pq[0] }

func main() {
	fmt.Println(maxPoints([][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, []int{5, 6, 2}))

}
