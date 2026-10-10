package main

import (
	"container/heap"
	"fmt"
	"math"
)

func openLock(deadends []string, target string) int {
	deMp := make(map[int]bool)
	for _, de := range deadends {
		num := 0
		for _, r := range de {
			num = num*10 + int(r-'0')
		}
		deMp[num] = true
	}
	tg := 0

	for _, r := range target {
		tg = tg*10 + int(r-'0')
	}

	G := make([][]int, 10000)
	for u := range 10000 {
		if deMp[u] {
			continue
		}
		for j := 1; j <= 1000; j *= 10 {
			v := u + j
			if (u/j)%10 == 9 {
				v = u - 9*j
			}
			if deMp[v] {
				continue
			}
			G[u] = append(G[u], v)
			G[v] = append(G[v], u)
		}
	}
	return dijkstra(G, tg)
}

func dijkstra(G [][]int, target int) int {
	dist := make([]int, 10000)
	for i := range 10000 {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	pq := make(PriorityQueue, 0)
	pq = append(pq, &Item{u: 0, dist: 0})
	heap.Init(&pq)
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u, d := it.u, it.dist
		if d > dist[u] {
			continue
		}

		for _, v := range G[u] {
			if dist[u]+1 < dist[v] {
				dist[v] = dist[u] + 1
				heap.Push(&pq, &Item{u: v, dist: dist[v]})
			}
		}
	}
	if dist[target] == math.MaxInt32 {
		return -1
	}
	return dist[target]
}

type Item struct {
	u    int // The value of the item; arbitrary.
	dist int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
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
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
