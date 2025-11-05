package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "container/heap"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// An IntHeap is a min-heap of ints.
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() int64 {
	heap := *h
	return heap[0]
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	for len(queue) > 0 {
		sum := 0
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		heap.Push(minHeap, sum)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	if k > minHeap.Len() {
		return -1
	}
	return minHeap.Peek()
}
