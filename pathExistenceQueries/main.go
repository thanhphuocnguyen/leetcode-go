package main

import "fmt"

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	parent := make([]int, n)
	weight := make([]int, n)
	for i := range n {
		parent[i] = i
		weight[i] = 1
	}

	var find func(u int) int
	find = func(u int) int {
		if parent[u] != u {
			parent[u] = find(parent[u])
		}
		return parent[u]
	}

	union := func(u, v int) {
		pu, pv := find(u), find(v)
		if pu == pv {
			return
		}
		if weight[pu] > weight[pv] {
			parent[pu] = pv
			weight[pv] += weight[pu]
		} else {
			parent[pv] = pu
			weight[pu] += weight[pv]
		}
	}

	for u, v := 0, 1; v < n; u, v = u+1, v+1 {
		if nums[v]-nums[u] <= maxDiff {
			union(u, v)
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = find(q[0]) == find(q[1])
	}

	return ans
}

func main() {
	fmt.Println(pathExistenceQueries(4, []int{2, 5, 6, 8}, 2, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}))
	fmt.Println(pathExistenceQueries(2, []int{1, 3}, 1, [][]int{{0, 0}, {0, 1}}))
}
