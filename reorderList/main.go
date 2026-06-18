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

func reorderList(head *ListNode) {
	// find mid using fast slow ptr
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	var firstPrev *ListNode
	for fast != nil && fast.Next != nil {
		firstPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// cut half list
	firstPrev.Next = nil
	// reverse from mid to end (prev is head of reversed from mid)
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// merge two list
	tmpHead := &ListNode{Val: 0}
	tmp := tmpHead
	for head != nil && prev != nil {
		tmp.Next = head
		head = head.Next
		tmp = tmp.Next
		tmp.Next = prev
		prev = prev.Next
		tmp = tmp.Next
	}
	for head != nil {
		tmp.Next = head
		tmp = tmp.Next
		head = head.Next
	}

	for prev != nil {
		tmp.Next = prev
		tmp = tmp.Next
		prev = prev.Next
	}
	head = tmpHead.Next
}

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reorderList(l)
	fmt.Println(l)
}
