package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	G := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if G[u] == nil {
			G[u] = []int{}
		}
		if G[v] == nil {
			G[v] = []int{}
		}
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	visited := make([]bool, n)
	var dfs func(s int) bool
	dfs = func(s int) bool {
		visited[s] = true
		if s == destination {
			return true
		}
		found := false
		for _, nei := range G[s] {
			if !visited[nei] {
				found = found || dfs(nei)
			}
		}

		return found
	}
	return dfs(source)
}
func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
}
