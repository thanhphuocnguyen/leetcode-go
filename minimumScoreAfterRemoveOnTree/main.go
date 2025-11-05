package main

import (
	"fmt"
	"math"
)

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	subTreeXor := make([]int, n)
	descendants := make([]map[int]bool, n)
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
	dfs = func(node, p int) {
		descendants[node][node] = true
		subTreeXor[node] = nums[node]
		for _, nei := range graph[node] {
			if p != nei {
				dfs(nei, node)
				subTreeXor[node] ^= nums[nei]
				descendants[node][nei] = true
			}
		}
	}
	dfs(0, -1)
	ans := math.MaxInt32
	totalXor := subTreeXor[0]
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xorI, xorJ := subTreeXor[i], subTreeXor[j]
			valI, valJ, valIJ := 0, 0, xorI^xorJ
			if _, ok := descendants[i][j]; ok {
				valJ = xorJ
				valI = totalXor ^ xorI
			} else if _, ok := descendants[j][i]; ok {
				valI = xorI
				valJ = totalXor ^ xorI
			} else {
				valI = xorI
				valJ = xorJ
				valIJ = totalXor ^ xorI ^ xorJ
			}
			mini := min(valI, min(valJ, valIJ))
			maxi := max(valI, max(valJ, valIJ))
			ans = min(ans, maxi-mini)
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumScore([]int{1, 5, 5, 4, 11}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
}
