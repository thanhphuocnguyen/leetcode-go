package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	visited := make(map[int]bool)
	graph := make(map[int][]int)
	ans := make([]int, 0)
	for _, adj := range adjacentPairs {
		u, v := adj[0], adj[1]
		if _, ok := graph[v]; !ok {
			graph[v] = make([]int, 0)
		}
		graph[v] = append(graph[v], u)
		graph[u] = append(graph[u], v)
	}
	node := -1
	for nd, nei := range graph {
		if len(nei) == 1 {
			node = nd
			break
		}
	}
	dfs(graph, visited, node, &ans)

	return ans
}

func dfs(graph map[int][]int, visited map[int]bool, node int, ans *[]int) {
	*ans = append(*ans, node)
	visited[node] = true
	for _, nei := range graph[node] {
		if !visited[nei] {
			dfs(graph, visited, nei, ans)
		}
	}
}

func main() {
	fmt.Println(restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}}))
	fmt.Println(restoreArray([][]int{{4, -2}, {1, 4}, {-3, 1}}))
}
