package main

import (
	"container/heap"
	"fmt"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	G := make([][][2]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([][2]int, 0)
	}
	for _, time := range times {
		u, v, w := time[0]-1, time[1]-1, time[2]
		G[u] = append(G[u], [2]int{v, w})
	}
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k-1] = 0
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{
		u: k - 1,
		d: 0,
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		u, d := node.u, node.d
		if d > dist[u] {
			continue
		}
		for _, nei := range G[u] {
			v, w := nei[0], nei[1]
			if w+dist[u] < dist[v] {
				dist[v] = w + dist[u]
				heap.Push(&pq, &Node{
					u: v,
					d: dist[v],
				})
			}
		}
	}
	ans := 0
	for _, d := range dist {
		if d == math.MaxInt32 {
			return -1
		}
		ans = max(ans, d)
	}
	return ans
}

// An Node is something we manage in a priority queue.
type Node struct {
	u int // The value of the item; arbitrary.
	d int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].d < pq[j].d
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
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
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
}
