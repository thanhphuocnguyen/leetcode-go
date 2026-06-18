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

func sortList(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

func mergeSort(left, right *ListNode) *ListNode {
	if left == nil || left == right {
		return left
	}
	prev, slow, fast := left, left, left
	for fast != right && fast.Next != right {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := prev.Next
	prev.Next = nil
	l1 := mergeSort(left, prev)
	l2 := mergeSort(next, right)
	return merge(l1, l2)
}

func merge(left, right *ListNode) *ListNode {
	tempHead := &ListNode{Val: 0}
	temp := tempHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			temp.Next = left
			left = left.Next
		} else {
			temp.Next = right
			right = right.Next
		}
		temp = temp.Next
	}

	for left != nil {
		temp.Next = left
		left = left.Next
		temp = temp.Next
	}

	for right != nil {
		temp.Next = right
		right = right.Next
		temp = temp.Next
	}

	return tempHead.Next
}

func main() {
	l := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	fmt.Println(sortList(l))
}
