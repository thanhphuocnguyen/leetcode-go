package main

import (
	"fmt"
	"math"
)

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	subTreeXor := make([]int, n)
	descendants := make([]map[int]bool, n)
	// build graph from edges
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, 0)
		descendants[i] = map[int]bool{}
	}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs func(node, p int)
	// dfs to calculate xor of all nodes
	dfs = func(node, p int) {
		// put node as its own descendant
		descendants[node][node] = true
		subTreeXor[node] = nums[node]
		for _, nei := range graph[node] {
			if p != nei {
				dfs(nei, node)
				// update subtree xor value
				subTreeXor[node] ^= nums[nei]
				// put neighbor in to node map as subtree
				descendants[node][nei] = true
			}
		}
	}
	dfs(0, -1)
	ans := math.MaxInt32
	// total xor is the xor value run from 0 by dfs
	totalXor := subTreeXor[0]
	// loop through all pair of nodes
	for a := 1; a < n; a++ {
		for b := a + 1; b < n; b++ {
			// take xor of a subtree and b subtree
			xorRootToA, xorRootToB := subTreeXor[a], subTreeXor[b]
			//
			compA, compB, compC := 0, 0, 0
			// if b/a is descendent of a/b. that mean a/b is in the same branch
			if _, ok := descendants[a][b]; ok {
				// xor subtree from b will be precomputed value of b
				compB = xorRootToB
				// xor subtree of a will be total xor value xor with xor of a subtree
				compA = totalXor ^ xorRootToA
				// xor sub tree of c is the the xor result of xor subtree from a xor with xor subtree from b
				compC = xorRootToA ^ xorRootToB
			} else if _, ok := descendants[b][a]; ok {
				// same as a when a is descendent of b
				compA = xorRootToA
				compB = totalXor ^ xorRootToA
				compC = xorRootToA ^ xorRootToB
			} else {
				// a and b in in different branches
				compA = xorRootToA
				compB = xorRootToB
				// c eq total ^ a ^ b
				compC = totalXor ^ xorRootToA ^ xorRootToB
			}
			mini := min(compA, min(compB, compC))
			maxi := max(compA, max(compB, compC))
			ans = min(ans, maxi-mini)
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumScore([]int{1, 5, 5, 4, 11}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
}
