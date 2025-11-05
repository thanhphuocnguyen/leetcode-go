package main

import "container/heap"

// An Item is something we manage in a priority queue.
type Cell struct {
	Row    int
	Col    int
	Height int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Height < pq[j].Height
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Cell)
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

func trapRainWater(heightMap [][]int) int {
	numOfRows, numOfCols := len(heightMap), len(heightMap[0])
	visited := make([][]bool, numOfRows)
	rowDir := []int{0, 0, -1, 1}
	colDir := []int{-1, 1, 0, 0}
	for i := range visited {
		visited[i] = make([]bool, numOfCols)
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	ans := 0
	for i := range numOfRows {
		// first col
		heap.Push(&pq, &Cell{Row: i, Col: 0, Height: heightMap[i][0]})
		visited[i][0] = true
		// last col
		heap.Push(
			&pq,
			&Cell{
				Row:    i,
				Col:    numOfCols - 1,
				Height: heightMap[i][numOfCols-1],
			})
		visited[i][numOfCols-1] = true
	}

	for i := range numOfCols {
		// first row
		heap.Push(&pq, &Cell{Row: 0, Col: i, Height: heightMap[0][i]})
		visited[0][i] = true
		// last row
		heap.Push(
			&pq,
			&Cell{
				Row:    numOfRows - 1,
				Col:    i,
				Height: heightMap[numOfRows-1][i],
			})
		visited[numOfRows-1][i] = true
	}
	for pq.Len() > 0 {
		currCell := heap.Pop(&pq).(*Cell)
		row, col, height := currCell.Row, currCell.Col, currCell.Height
		for i := 0; i < 4; i++ {
			nextRow := row + rowDir[i]
			nextCol := col + colDir[i]
			if isValid(nextRow, nextCol, numOfRows, numOfCols) && !visited[nextRow][nextCol] {
				neighborHeight := heightMap[nextRow][nextCol]
				if neighborHeight < height {
					ans += height - neighborHeight
				}
				heap.Push(
					&pq,
					&Cell{
						Row:    nextRow,
						Col:    nextCol,
						Height: max(neighborHeight, height),
					})
				visited[nextRow][nextCol] = true
			}
		}
	}
	return ans
}

func isValid(row, col, numOfRows, numOfCols int) bool {
	return row >= 0 && col >= 0 && row < numOfRows && col < numOfCols
}
