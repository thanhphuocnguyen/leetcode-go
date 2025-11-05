package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	temp := head
	for temp != nil && temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head
}

func main() {
	vals := []int{2, 6, 3, 4, 5, 6}
	head := &ListNode{
		Val: 1,
	}
	temp := head
	for _, val := range vals {
		temp.Next = &ListNode{
			Val: val,
		}
		temp = temp.Next
	}
	fmt.Println(removeElements(head, 6))
}
