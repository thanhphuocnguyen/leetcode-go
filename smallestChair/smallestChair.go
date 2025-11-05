package main

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Event [2]int

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Event

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Event)
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

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	arrivals := make([]int, n)
	for i := 0; i < n; i++ {
		arrivals[i] = i
	}

	sort.Slice(arrivals, func(i, j int) bool {
		return times[arrivals[i]][0] < times[arrivals[j]][0]
	})

	emptySeats := &IntHeap{}
	heap.Init(emptySeats)
	for i := 0; i < n; i++ {
		heap.Push(emptySeats, i)
	}

	occupiedSeats := &PriorityQueue{}
	heap.Init(occupiedSeats)

	for _, order := range arrivals {
		arrival, leave := times[arrivals[order]][0], times[arrivals[order]][1]

		for occupiedSeats.Len() > 0 && (*occupiedSeats)[0][0] <= arrival {
			event := heap.Pop(occupiedSeats).(*Event)
			heap.Push(emptySeats, event[1])
		}

		seat := heap.Pop(emptySeats).(int)
		if arrivals[order] == targetFriend {
			return seat
		}

		heap.Push(occupiedSeats, &Event{leave, seat})
	}

	return -1
}
