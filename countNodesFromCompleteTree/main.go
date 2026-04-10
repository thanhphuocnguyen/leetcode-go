package main

import "fmt"

/**
 * Definition for a binary tree node.
 *
 * TreeNode represents a single node in the binary tree
 * - Val: the value stored in this node
 * - Left: pointer to the left child node (nil if no left child)
 * - Right: pointer to the right child node (nil if no right child)
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// countNodes counts the total number of nodes in a complete binary tree.
//
// Algorithm:
// For a complete binary tree, we compare the height of the leftmost path of
// the left subtree with the height of the leftmost path of the right subtree.
//
// Case 1: If leftHeight == rightHeight
//   - The left subtree is a perfect binary tree (all levels fully filled)
//   - Number of nodes in left subtree = 2^leftHeight - 1
//   - Recursively count nodes in the right subtree
//   - Total = (1 << leftHeight) + countNodes(right)
//
// Case 2: If leftHeight > rightHeight
//   - The right subtree is a perfect binary tree
//   - Number of nodes in right subtree = 2^rightHeight - 1
//   - Recursively count nodes in the left subtree
//   - Total = (1 << rightHeight) + countNodes(left)
//
// Time Complexity: O((log n)^2) because height calculation is O(log n)
// and we do this at most O(log n) times
// Space Complexity: O(log n) for recursion stack
func countNodes(root *TreeNode) int {
	// Base case: empty tree has 0 nodes
	if root == nil {
		return 0
	}

	// Get the height by traversing the leftmost path of both subtrees
	leftLH := getLefthHeight(root.Left)
	rightLH := getLefthHeight(root.Right)

	// Compare the two heights to determine which subtree is perfect
	if leftLH == rightLH {
		// Left subtree is a perfect binary tree with height leftLH
		// So left subtree has (1 << leftLH) nodes (which is 2^leftLH)
		// Add 1 for the current root, then recursively count right subtree
		return (1 << leftLH) + countNodes(root.Right)
	} else {
		// Right subtree is a perfect binary tree with height rightLH
		// So right subtree has (1 << rightLH) nodes (which is 2^rightLH)
		// Add 1 for the current root, then recursively count left subtree
		return (1 << rightLH) + countNodes(root.Left)
	}
}

// getLefthHeight calculates the height of the binary tree by traversing
// the leftmost path (going left as much as possible).
//
// This is useful for complete binary trees because:
// - The leftmost path always represents the actual height of the tree
// - This helps us determine if subtrees are perfect binary trees
//
// Returns: the height (number of nodes from root to leaf on leftmost path)
func getLefthHeight(root *TreeNode) int {
	rs := 0
	node := root

	// Traverse left until we reach a nil node
	for node != nil {
		rs++
		node = node.Left
	}

	// Return the count of nodes on the leftmost path
	return rs
}

// main demonstrates the countNodes function with a sample complete binary tree
func main() {
	// Create a complete binary tree:
	//       1
	//      / \
	//     2   3
	//    / \ /
	//   4  5 6
	// This tree has 6 nodes total
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 6}}}

	// Count and print the total number of nodes
	fmt.Println((countNodes(root)))
}
