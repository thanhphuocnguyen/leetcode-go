package main

import "container/heap"

func processQueries(c int, connections [][]int, queries [][]int) []int {
	// build graph
	graph := make([][]int, c+1)
	for i := 1; i <= c; i++ {
		graph[i] = make([]int, 0)
	}
	offlines := make([]bool, c+1)
	stationIds := make([]int, c+1)
	for i := range c + 1 {
		stationIds[i] = -1
	}
	// add adj
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs func(vtx int, pId int, h *IntHeap)
	dfs = func(vtx int, pId int, h *IntHeap) {
		stationIds[vtx] = pId
		heap.Push(h, vtx)
		for _, nei := range graph[vtx] {
			if stationIds[nei] == -1 {
				dfs(nei, pId, h)
			}
		}
	}
	// create heapMap for each ID
	heapMp := make([]*IntHeap, 0)
	for i, id := 1, 0; i <= c; i++ {
		if len(graph[i]) == 0 {
			continue
		}
		v := graph[i][0]
		// try to dfs each station and init min heap for each traversal
		if stationIds[v] == -1 {
			h := &IntHeap{}
			dfs(v, id, h)
			heap.Init(h)
			heapMp = append(heapMp, h)
			id++
		}
	}
	ans := make([]int, 0, len(queries))
	for _, q := range queries {
		typ, stat := q[0], q[1]
		if typ == 2 {
			offlines[stat] = true
		} else {
			if !offlines[stat] {
				ans = append(ans, stat)
			} else {
				// heapMp
				pId := stationIds[stat]
				if pId == -1 {
					ans = append(ans, -1)
					continue
				}
				for heapMp[pId].Len() > 0 && offlines[heapMp[pId].Peek()] {
					heap.Pop(heapMp[pId])
				}
				if heapMp[pId].Len() > 0 {
					ans = append(ans, heapMp[pId].Peek())
				} else {
					ans = append(ans, -1)
				}
			}
		}
	}
	return ans
}

type IntHeap []int

// Len implements heap.Interface.
func (h IntHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h IntHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements heap.Interface.
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Push implements heap.Interface.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h IntHeap) Peek() int {
	return h[0]
}

// Swap implements heap.Interface.
func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	processQueries(3, [][]int{}, [][]int{{1, 1}, {2, 1}, {1, 1}})
	processQueries(3, [][]int{{2, 3}, {1, 2}, {1, 3}}, [][]int{{1, 1}, {1, 2}, {1, 2}, {1, 3}, {1, 3}, {1, 1}, {2, 3}, {1, 1}, {2, 2}, {2, 2}, {1, 2}, {1, 3}, {2, 1}, {2, 1}, {1, 3}, {2, 1}, {2, 3}, {1, 3}, {1, 3}, {2, 2}, {1, 1}, {2, 2}, {1, 2}, {1, 1}, {1, 2}, {1, 3}, {1, 2}, {1, 3}, {2, 2}, {2, 2}, {2, 3}, {1, 3}, {1, 2}, {2, 3}, {1, 2}, {2, 3}, {2, 3}, {2, 2}, {2, 2}, {1, 1}, {2, 3}, {1, 1}})
	processQueries(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}})
	processQueries(2, [][]int{{2, 1}}, [][]int{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {2, 2}, {2, 2}, {2, 1}, {1, 1}, {2, 2}, {2, 1}, {1, 1}, {2, 1}, {2, 2}, {1, 1}, {1, 2}, {1, 1}, {1, 1}, {2, 1}, {1, 1}, {2, 2}, {2, 2}, {2, 1}, {1, 2}, {2, 2}, {1, 1}, {2, 1}, {2, 2}, {2, 2}, {1, 1}, {2, 2}})
}

// [{1,1},{1,2},{1,2},{1,3},{1,3},{1,1},{2,3},{1,1},{2,2},{2,2},{1,2},{1,3},{2,1},{2,1},{1,3},{2,1},{2,3},{1,3},{1,3},{2,2},{1,1},{2,2},{1,2},{1,1},{1,2},{1,3},{1,2},{1,3},{2,2},{2,2},{2,3},{1,3},{1,2},{2,3},{1,2},{2,3},{2,3},{2,2},{2,2},{1,1},{2,3},{1,1}]
