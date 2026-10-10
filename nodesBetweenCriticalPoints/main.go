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

func nodesBetweenCriticalPoints(head *ListNode) []int {
	ans := []int{-1, -1}
	if head == nil {
		return ans
	}
	var prev, curr *ListNode = head, head.Next
	firstPos, prevPos := -1, -1
	i := 1
	for curr != nil {
		if check(curr, curr.Next, prev) {
			if firstPos == -1 {
				firstPos = i
				prevPos = i
			} else {
				if prevPos != -1 {
					if ans[0] == -1 {
						ans[0] = i - prevPos
					} else {
						ans[0] = min(i-prevPos, ans[0])

					}
					ans[1] = (i - firstPos)
				}
				prevPos = i
			}
		}
		i++
		prev = curr
		curr = curr.Next
	}

	return ans
}

func check(curr, next, prev *ListNode) bool {
	if next != nil && prev != nil {
		return (curr.Val > next.Val && curr.Val > prev.Val) || (curr.Val < next.Val && curr.Val < prev.Val)
	}
	return false
}

func main() {
	var list *ListNode = &ListNode{}
	tc1 := []int{1, 3, 2, 2, 3, 2, 2, 2, 7}
	l1 := list
	for i, num := range tc1 {
		list.Val = num
		if i != len(tc1)-1 {
			list.Next = &ListNode{}
		}
		list = list.Next
	}
	fmt.Println(nodesBetweenCriticalPoints(l1))
}
