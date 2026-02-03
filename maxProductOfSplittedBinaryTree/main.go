package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MOD = 1_000_000_007

func maxProduct(root *TreeNode) int {
	totalSum := getSum(root)

	var ans int = 0
	getMaxProduct(root, totalSum, &ans)
	return ans % MOD
}

func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val + getSum(root.Left) + getSum(root.Right)
	return ans
}

func getMaxProduct(root *TreeNode, total int, ans *int) int {
	if root == nil {
		return 0
	}

	sum := root.Val + getMaxProduct(root.Left, total, ans) + getMaxProduct(root.Right, total, ans)

	*ans = max(*ans, (total-sum)*sum)
	return sum
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(maxProduct(root))
}
