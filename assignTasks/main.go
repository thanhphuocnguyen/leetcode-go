package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	first  int
	second int
}

type PriorityQueue struct {
	data []Item
	less func(a, b Item) bool
}

func assignTasks(servers []int, tasks []int) []int {
	busy := &PriorityQueue{
		less: func(a, b Item) bool {
			return a.first < b.first
		},
	}
	idle := &PriorityQueue{
		less: func(a, b Item) bool {
			if a.first != b.first {
				return a.first < b.first
			}
			return a.second < b.second
		},
		data: make([]Item, len(servers)),
	}
	for i, w := range servers {
		idle.data[i] = Item{first: w, second: i}
	}
	heap.Init(idle)
	heap.Init(busy)
	ts := 0
	ans := make([]int, len(tasks))
	release := func() {
		for busy.Len() > 0 && busy.Top().first <= ts {
			s := heap.Pop(busy).(Item)
			heap.Push(idle, Item{servers[s.second], s.second})
		}
	}
	for i, time := range tasks {
		ts = max(ts, time)
		release()
		if idle.Len() == 0 {
			ts = busy.Top().first
			release()
		}

		s := heap.Pop(idle).(Item)
		ans[i] = s.second
		heap.Push(busy, Item{ts + time, s.second})
	}

	return ans
}

// Len implements [heap.Interface].
func (p PriorityQueue) Len() int {
	return len(p.data)
}

// Less implements [heap.Interface].
func (p PriorityQueue) Less(i int, j int) bool {
	return p.less(p.data[i], p.data[j])
}

// Pop implements [heap.Interface].
func (p *PriorityQueue) Pop() any {
	old := p.data
	n := len(old)
	it := old[n-1]

	p.data = old[:n-1]
	return it
}

func (h *PriorityQueue) Top() Item { return h.data[0] }

// Push implements [heap.Interface].
func (p *PriorityQueue) Push(x any) {
	item := x.(Item)
	p.data = append(p.data, item)
}

// Swap implements [heap.Interface].
func (p PriorityQueue) Swap(i int, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func main() {
	fmt.Println(assignTasks([]int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}))
	fmt.Println(assignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}))
}
