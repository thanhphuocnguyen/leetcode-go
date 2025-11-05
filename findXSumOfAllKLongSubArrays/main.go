package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].value > pq[j].value
	}
	return pq[i].priority > pq[j].priority
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
	old[n-1] = nil // Avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	freq := make([]int, 51)
	ans := make([]int, n-k+1)
	aIdx := 0
	for i, j := 0, 0; j < len(nums); {
		left, right := nums[i], nums[j]
		freq[right]++
		if j-i+1 == k {
			pq := make(PriorityQueue, 0)
			for i, f := range freq {
				if f > 0 {
					heap.Push(&pq, &Item{
						value:    i,
						priority: f,
					})
				}
			}
			sum := 0
			pqLen := pq.Len()
			for i := 0; i < min(x, pqLen); i++ {
				peek := heap.Pop(&pq).(*Item)
				sum += peek.value * peek.priority
			}
			ans[aIdx] = sum
			freq[left]--
			i++
			aIdx++
		}
		j++
	}
	return ans
}

func main() {
	fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
}
