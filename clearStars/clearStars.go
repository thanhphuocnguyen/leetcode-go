package main

import (
	"container/heap"
	"strings"
)

func clearStars(s string) string {
	pq := make(PriorityQueue, 0)
	var b strings.Builder
	removedItems := make([]bool, len(s))
	for i := range s {
		c := s[i]
		if c == '*' {
			removedItems[i] = true
			if pq.Len() > 0 {
				item := heap.Pop(&pq).(*Item)
				removedItems[item.idx] = true
			}
		} else {
			heap.Push(&pq, &Item{i, s[i]})
		}
	}
	for i, removed := range removedItems {
		if !removed && s[i] != '*' {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

type Item struct {
	idx      int  // The value of the item; arbitrary.
	priority byte // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not lowest, priority so we use greater than here.
	if pq[i].priority == pq[j].priority {
		return pq[i].idx > pq[j].idx
	}
	return pq[i].priority < pq[j].priority
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
	println(clearStars("aaba*"))
	println(clearStars("abc"))
	println(clearStars("yd**"))
}
