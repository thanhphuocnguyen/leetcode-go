package main

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr, next := head, head.Next
	var prev *ListNode
	tempNext := next

	for curr != nil && curr.Next != nil {
		curr.Next = tempNext.Next
		tempNext.Next = curr
		if prev != nil {
			prev.Next = tempNext
		}
		prev = curr
		curr = curr.Next
		if curr == nil {
			break
		}
		tempNext = curr.Next
	}

	return next
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	swapPairs(list)
	print(list)
}
