package main

import "container/heap"

func lastStoneWeight(stones []int) int {
	h := &IntMaxHeap{}
	*h = append(*h, stones...)
	heap.Init(h) // Initialize the heap
	for h.Len() > 1 {
		heap.Push(h, heap.Pop(h).(int)-heap.Pop(h).(int))
	}

	if h.Len() > 0 {
		return h.Pop().(int)
	}

	return 0
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int { return len(h) }

// For a max heap, Less should return true if h[i] is *less* than h[j]
// so that the heap maintains the property where parents are greater than children.
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
