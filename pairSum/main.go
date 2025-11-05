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

func pairSum(head *ListNode) int {
	var vals = [100000]int{}
	i := 0
	for head != nil {
		vals[i] = head.Val
		i++
		head = head.Next
	}
	ans := 0
	for j := 0; j < i/2; j++ {
		ans = max(ans, vals[j]+vals[i-j-1])
	}
	return ans
}
