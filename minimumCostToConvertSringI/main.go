package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumCost("abcd", "acbe", []byte{'a', 'b', 'c', 'c', 'e', 'd'}, []byte{'b', 'c', 'b', 'e', 'b', 'e'}, []int{2, 5, 5, 1, 2, 20}))
	fmt.Println(minimumCost("aaaa", "bbbb", []byte{'a', 'c'}, []byte{'c', 'b'}, []int{1, 2}))
	fmt.Println(minimumCost("abcd", "abce", []byte{'a'}, []byte{'e'}, []int{10000}))
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var dist [26][26]int
	for i := range 26 {
		for j := range 26 {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	for i := range original {
		a, b, c := original[i]-'a', changed[i]-'a', cost[i]
		dist[a][b] = min(dist[a][b], c)
	}

	floyWarshall(&dist)

	var ans int64
	for i := range source {
		d := dist[source[i]-'a'][target[i]-'a']
		if d == math.MaxInt32 {
			return -1
		}
		ans += int64(d)
	}

	return ans
}

func floyWarshall(dist *[26][26]int) {
	// k is the bridge from i -> k -> j find the min value dist each path
	for k := range 26 {
		for i := range 26 {
			for j := range 26 {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}
}

// func dijkstra(adjacent map[byte]map[byte]int, src, dest byte) int {
// 	pq := PriorityQueue{&Item{
// 		Node: src,
// 		Dist: 0,
// 	}}
// 	dist := make(map[byte]int)

// 	dist[src] = 0
// 	heap.Init(&pq)
// 	for pq.Len() > 0 {
// 		it := heap.Pop(&pq).(*Item)
// 		u := it.Node
// 		for v, c := range adjacent[u] {
// 			if _, ok := dist[v]; !ok {
// 				dist[v] = math.MaxInt32
// 			}
// 			if dist[u]+c < dist[v] {
// 				dist[v] = dist[u] + c
// 				heap.Push(&pq, &Item{Node: v, Dist: dist[v]})
// 			}
// 		}
// 	}
// 	return dist[dest]
// }

// type Item struct {
// 	Node byte // The value of the item; arbitrary.
// 	Dist int  // The priority of the item in the queue.
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
// 	return pq[i].Dist < pq[j].Dist
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// func (pq *PriorityQueue) Push(x any) {
// 	item := x.(*Item)
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
// 	*pq = old[0 : n-1]
// 	return item
// }
