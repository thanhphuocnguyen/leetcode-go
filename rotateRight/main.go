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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := 0
	temp := head
	for temp != nil {
		n++
		temp = temp.Next
	}
	k = k % n
	if k == 0 {
		return head
	}

	var prev *ListNode
	temp = head
	for i := 0; i < n-k; i++ {
		prev = temp
		temp = temp.Next
	}
	if prev != nil && prev.Next != nil {
		prev.Next = nil
	}
	newHead := temp
	temNewHead := temp
	for temNewHead != nil && temNewHead.Next != nil {
		temNewHead = temNewHead.Next
	}
	temNewHead.Next = head
	return newHead
}

func main() {
	// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Println(rotateRight(head1, 1))
}
