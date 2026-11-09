package main

import (
	"container/heap"
	"fmt"
	"math"
)

func shortestPath(n int, edges [][]int, labels string, k int) int {
	graph := make([][][2]int, n)
	for i := range n {
		graph[i] = make([][2]int, 0)
	}
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], [2]int{v, w})
	}

	return dijkstra(graph, n, k, labels)
}

type Item struct {
	u int
	d int
	f int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].d == pq[j].d {
		return pq[i].f < pq[j].f
	}
	return pq[i].d < pq[j].d
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	// n := len(*pq)
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func dijkstra(graph [][][2]int, n, k int, labels string) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][1] = 0

	pq := PriorityQueue{{u: 0, d: 0, f: 1}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)

		u, d, f := it.u, it.d, it.f

		if d > dist[u][f] {
			continue
		}

		if u == n-1 {
			return d
		}

		for _, nei := range graph[u] {
			v, w := nei[0], nei[1]
			newD := d + w
			newF := 1
			if labels[v] == labels[u] {
				newF = f + 1
			}
			if newD < dist[v][newF] {

				if newF <= k {
					heap.Push(&pq, &Item{u: v, d: newD, f: newF})
					dist[v][newF] = newD
				}
			}

		}
	}
	return -1
}

func main() {
	fmt.Println(shortestPath(3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}, "aab", 2))
	fmt.Println(shortestPath(3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}, "aab", 1))
}
