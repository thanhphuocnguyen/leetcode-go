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

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	return buildTree(n, map[int][]*TreeNode{})
}

func buildTree(n int, memo map[int][]*TreeNode) []*TreeNode {
	if n == 1 {
		return []*TreeNode{{}}
	}
	if val, ok := memo[n]; ok {
		return val
	}
	ans := make([]*TreeNode, 0)
	for i := 1; i < n; i += 2 {
		leftNodes := buildTree(i, memo)
		rightNodes := buildTree(n-i-1, memo)
		for _, lNode := range leftNodes {
			for _, rNode := range rightNodes {
				ans = append(ans, &TreeNode{Left: lNode, Right: rNode})
			}
		}
	}
	memo[n] = ans
	return memo[n]
}

func main() {
	fmt.Println(allPossibleFBT(5))
}
