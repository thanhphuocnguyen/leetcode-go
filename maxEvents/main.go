package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	h := &IntHeap{}
	heap.Init(h)
	ans := 0
	i := 0
	occupiedDay := 1
	for h.Len() > 0 || i < len(events) {
		// add all events have start day less than occupied day
		for i < len(events) && events[i][0] <= occupiedDay {
			heap.Push(h, events[i][1])
			i++
		}
		// remove all events have end day less than occupied day
		for h.Len() > 0 && (*h)[0] < occupiedDay {
			heap.Pop(h)
		}

		// use smallest event participated
		if h.Len() > 0 {
			heap.Pop(h)
			ans++
		}

		// increase day
		occupiedDay++
	}

	return ans
}

func main() {
	// fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}))
	fmt.Println(maxEvents([][]int{{1, 2}, {1, 2}, {1, 5}, {1, 5}, {3, 3}}))
}
