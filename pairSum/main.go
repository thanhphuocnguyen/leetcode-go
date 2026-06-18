package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// find mid of ll
	// reverse from mid to tail
	// find max sum of head + revHead from til n/2
	fast, slow := head, head
	n := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		n++
	}

	mid := reverse(slow)
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, head.Val+mid.Val)
		head = head.Next
		mid = mid.Next
	}
	return ans
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	curr := node
	// 1 -> 2 -> 3
	// 1 next = nil, then 2 2.next = 1->nil, prev = curr
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	list := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	fmt.Println(pairSum(list))
}
