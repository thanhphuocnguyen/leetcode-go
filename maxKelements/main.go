package main

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func maxKelements(nums []int, k int) int64 {
	maxHeap := make(MaxHeap, len(nums))
	heap.Init(&maxHeap)
	for _, num := range nums {
		heap.Push(&maxHeap, num)
	}
	var score int64
	for i := 0; i < k; i++ {
		num := (heap.Pop(&maxHeap)).(int64)
		score += int64(num)
		heap.Push(&maxHeap, int(math.Ceil((float64(num))/3)))
	}

	return score
}
