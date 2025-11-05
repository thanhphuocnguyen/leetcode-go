package main

import "container/heap"

type Item struct {
	val  int
	freq int
}
type PriorityQueue []*Item

// Swap implements heap.Interface.
func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Len implements heap.Interface.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less implements heap.Interface.
func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].freq > pq[j].freq
}

// Pop implements heap.Interface.
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := old.Len()
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

// Push implements heap.Interface.
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	pq := make(PriorityQueue, 0)
	for k, v := range mp {
		pq = append(pq, &Item{
			val:  k,
			freq: v,
		})
	}
	heap.Init(&pq)
	ans := make([]int, k)
	for i := 0; i < min(k, len(mp)); i++ {
		ans[i] = heap.Pop(&pq).(*Item).val
	}
	return ans
}
