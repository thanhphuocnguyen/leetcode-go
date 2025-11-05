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

func modifiedList(nums []int, head *ListNode) *ListNode {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	dummy := &ListNode{Val: -1, Next: head}
	curr := dummy
	for curr != nil && curr.Next != nil {
		if numSet[curr.Next.Val] {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

func main() {
	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	nums1 := []int{1}
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	fmt.Println(modifiedList(nums1, head1))
}
