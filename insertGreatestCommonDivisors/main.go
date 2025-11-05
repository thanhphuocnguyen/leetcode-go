package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	temp := head
	for node := temp; node.Next != nil; node = node.Next {
		nextNode := node.Next
		newNode := ListNode{
			Val:  gcd(node.Val, nextNode.Val),
			Next: nextNode,
		}
		temp.Next = &newNode
		temp = nextNode
	}

	return head
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
