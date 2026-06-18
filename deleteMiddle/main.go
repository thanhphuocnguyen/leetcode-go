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

func deleteMiddle(head *ListNode) *ListNode {
	slow, fast, prevSlow := head, head, head
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow, fast = slow.Next, fast.Next.Next
	}
	if slow == fast {
		return nil
	}
	if prevSlow.Next != nil {
		prevSlow.Next = prevSlow.Next.Next
	}
	return head
}

func main() {
	ll := &ListNode{Val: 1}
	// ll := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}}}}}
	fmt.Println(deleteMiddle(ll))
}
