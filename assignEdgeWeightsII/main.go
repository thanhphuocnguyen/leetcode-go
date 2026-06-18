package main

import "fmt"

const MOD = 1_000_000_007

const LOG = 17

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = (pow2[i-1] * 2) % MOD
	}
	parent := make([]int, n+1)
	depth := make([]int, n+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		parent[v] = u
	}
	for v := 1; v <= n; v++ {
		if parent[v] != 0 {
			depth[v] = depth[parent[v]] + 1
		}
	}

	// dynamic programming for kth ancestors precomputed
	up := make([][]int, n+1)
	for v := 0; v <= n; v++ {

		up[v] = make([]int, LOG)
		up[v][0] = parent[v]
	}

	for j := 1; j < LOG; j++ {
		for v := 1; v <= n; v++ {
			// 2^j = 2*(2^j-1) <=> 2^j-1 + 2^j-1
			up[v][j] = up[up[v][j-1]][j-1]
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]
		if u == v {
			ans[i] = 0
		} else {
			lca := lca(up, depth, u, v)
			dist := depth[u] + depth[v] - 2*(depth[lca])
			// The number of edges on the path from u to v is dist, and we can assign weights in 2^(dist-1) ways
			ans[i] = pow2[dist-1]
		}
	}
	return ans
}

func lca(up [][]int, depth []int, u, v int) int {
	if depth[u] < depth[v] {
		u, v = v, u
	}

	diff := depth[u] - depth[v]
	for j := LOG - 1; j >= 0; j-- {
		if ((diff >> j) & 1) == 1 {
			u = up[u][j]
		}
	}

	if u == v {
		return u
	}
	for j := LOG - 1; j >= 0; j-- {
		if up[u][j] != up[v][j] {
			u = up[u][j]
			v = up[v][j]
		}
	}

	return up[u][0]
}

func main() {
	fmt.Println(assignEdgeWeights([][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}, [][]int{{1, 4}, {3, 4}, {2, 5}}))
}
