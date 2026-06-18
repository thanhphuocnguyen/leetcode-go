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

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	var firstList, tmpFist, secondList, tmpSecond *ListNode
	for temp != nil {
		if temp.Val < x {
			if firstList == nil {
				firstList = temp
				tmpFist = temp
			} else {
				tmpFist.Next = temp
				tmpFist = tmpFist.Next
			}
		} else {
			if secondList == nil {
				secondList = temp
				tmpSecond = temp
			} else {
				tmpSecond.Next = temp
				tmpSecond = tmpSecond.Next
			}
		}
		temp = temp.Next
	}
	if firstList == nil {
		return secondList
	} else if secondList == nil {
		return firstList
	}

	tmpSecond.Next = nil
	tmpFist.Next = secondList

	return firstList
}

func main() {
	// list := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	// x := 3
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 1}}
	partition(list2, 0)
}
