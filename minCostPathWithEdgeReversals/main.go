package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCost(4, [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}}))
}

type Item struct {
	Node int
	Dist int
}

func minCost(n int, edges [][]int) int {
	graph := make([][][2]int, n)
	for i := range n {
		graph[i] = make([][2]int, 0)
	}
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if u != 0 {
			graph[u] = append(graph[u], [2]int{v, w})
			graph[v] = append(graph[v], [2]int{u, 2 * w})
		} else {
			graph[u] = append(graph[u], [2]int{v, w})
		}
	}
	return dijkstra(n, graph)
}

func dijkstra(n int, graph [][][2]int) int {
	pq := make(PriorityQueue, 0)
	pq = append(pq, &Item{Node: 0, Dist: 0})
	heap.Init(&pq)
	dist := make([]int, n)

	for i := 1; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.Node
		for _, nei := range graph[item.Node] {
			v, w := nei[0], nei[1]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(&pq, &Item{
					Node: v,
					Dist: dist[v],
				})
			}
		}
	}

	return dist[n-1]
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Dist < pq[j].Dist
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
