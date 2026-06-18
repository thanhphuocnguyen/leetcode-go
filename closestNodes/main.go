package main

import (
	"fmt"
)

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

func closestNodes(root *TreeNode, queries []int) [][]int {
	arr := make([]int, 0)
	inOrderTraversal(root, &arr)
	ans := make([][]int, len(queries))
	for i, q := range queries {
		// u, l := upperBound(arr, q), lowerBound(arr, q)
		ans[i] = upperBound(arr, q)
	}
	return ans
}

func upperBound(arr []int, target int) []int {
	lo, hi := 0, len(arr)-1
	lower, upper := -1, -1
	for lo <= hi {
		m := (lo + hi) / 2
		mid := arr[m]
		if mid == target {
			return []int{mid, mid}
		} else if mid > target {
			upper = mid
			hi = m - 1
		} else {
			lower = mid
			lo = m + 1
		}
	}

	return []int{lower, upper}
}
func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	lower := -1
	for lo <= hi {
		m := (lo + hi) / 2
		mid := arr[m]
		if mid == target {
			return mid
		} else if mid < target {
			lower = mid
			lo = m + 1
		} else {
			hi = m - 1
		}
	}

	return lower
}

func inOrderTraversal(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrderTraversal(root.Right, arr)
}

func main() {
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 13,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 14,
				},
			},
		},
	}
	fmt.Println(closestNodes(tree, []int{5}))
}
