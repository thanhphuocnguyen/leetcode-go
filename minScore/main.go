package main

import (
	"fmt"
	"math"
)

func minScore(n int, roads [][]int) int {
	G := make([][][2]int, n)
	for i := range n {
		G[i] = make([][2]int, 0)
	}
	for _, road := range roads {
		u, v, d := road[0]-1, road[1]-1, road[2]
		G[u] = append(G[u], [2]int{v, d})
		G[v] = append(G[v], [2]int{u, d})
	}
	visited := make([]bool, n)

	var dfs func(u int)
	ans := math.MaxInt32
	dfs = func(u int) {
		visited[u] = true

		for _, nei := range G[u] {
			v, d := nei[0], nei[1]
			if d < ans {
				ans = d
			}
			if !visited[nei[0]] {
				dfs(v)
			}
		}
	}

	dfs(0)
	return ans
}

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
}
