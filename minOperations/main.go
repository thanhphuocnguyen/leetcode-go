package main

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	value    int64 // The value of the item; arbitrary.
	priority int64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use greater than here.
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

func minOperations(nums []int, k int) int {
	pq := make(PriorityQueue, len(nums))
	for i, num := range nums {
		pq[i] = &Item{
			value:    int64(num),
			priority: int64(num),
		}
	}
	ans := 0
	heap.Init(&pq)
	for pq.Len() > 1 {
		x := heap.Pop(&pq).(*Item)
		if x.value >= int64(k) {
			break
		}
		y := heap.Pop(&pq).(*Item)
		op := min(x.value, y.value)*2 + max(x.value, y.value)
		heap.Push(&pq, &Item{
			value:    op,
			priority: op,
		})
		ans++
	}
	return ans
}

func main() {
	nums := []int{2, 11, 10, 1, 3}
	k := 10
	minOperations(nums, k)
}
