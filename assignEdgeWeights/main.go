package main

import (
	"fmt"
)

const MOD = 1_000_000_007

// Fermat little's theorem
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	p := power(x, y/2) % MOD
	p = (p * p) % MOD
	if y%2 == 0 {
		return p
	} else {
		return (x * p) % MOD
	}
}

func fastExp(base, exp int) int {
	res := 1
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		exp /= 2
		base = (base * base) % MOD
	}

	return res
}

func assignEdgeWeights(edges [][]int) int {
	tree := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if _, ok := tree[u]; !ok {
			tree[u] = make([]int, 0)
		}

		if _, ok := tree[v]; !ok {
			tree[v] = make([]int, 0)
		}

		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}
	visited := make([]bool, len(tree)+1)
	maxDepth := dfs(tree, visited, 1, 0)
	return fastExp(2, maxDepth-1)
}

func dfs(tree map[int][]int, visited []bool, src int, depth int) int {
	visited[src] = true
	currDepth := depth
	for _, child := range tree[src] {
		if visited[child] {
			continue
		}
		currDepth = max(currDepth, dfs(tree, visited, child, depth+1))
	}
	return currDepth
}

func main() {
	fmt.Println(assignEdgeWeights([][]int{{3, 2}, {2, 1}}))
	fmt.Println(assignEdgeWeights([][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}))
}
