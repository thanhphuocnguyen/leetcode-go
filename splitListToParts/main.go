package main

func main() {
	// Test cases
	// 1. [1,2,3], k = 5
	// 2. [1,2,3,4,5,6,7,8,9,10], k = 3
	// 3. [1,2,3,4,5,6,7,8,9,10], k = 7
	// 4. [1,2,3,4,5,6,7,8,9,10], k = 1
	// 5. [1,2,3,4,5,6,7,8,9,10], k = 10
	// 6. [1,2,3,4,5,6,7,8,9,10], k = 11
	// 7. [1,2,3,4,5,6,7,8,9,10], k = 12

	// Test case 1
	splitListToParts(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, 5)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	// Get the length of the list
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	// Calculate the number of elements in each part
	elementsPerPart := length / k
	remainder := length % k

	// Create the result
	result := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		// Create the part
		result[i] = head

		// Calculate the number of elements in this part
		elements := elementsPerPart
		if i < remainder {
			elements++
		}

		// Move to the next part
		for j := 0; j < elements; j++ {
			if j == elements-1 {
				// If this is the last element of this part, set the next element to nil
				head, head.Next = head.Next, nil
			} else {
				head = head.Next
			}
		}
	}

	return result
}
