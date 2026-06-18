package main

import "fmt"

/*
*
Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	kthStart, kthEnd := head, head
	n := 0
	for temp := head; temp != nil; temp = temp.Next {
		if n == k-1 {
			kthStart = temp
		}
		n++
	}

	for i := 0; i < n-k; i++ {
		kthEnd = kthEnd.Next
	}
	kthStart.Val, kthEnd.Val = kthEnd.Val, kthStart.Val
	return head
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	fmt.Println(swapNodes(list, 2))
}
