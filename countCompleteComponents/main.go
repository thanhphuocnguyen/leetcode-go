package main

func countCompleteComponents(n int, edges [][]int) int {
	graph := buildGraph(n, edges)
	visited := make([]bool, n)
	ans := 0
	for i := range n {
		if !visited[i] {
			edgeCnt := 0
			nodeCnt := 0
			dfs(graph, i, visited, &edgeCnt, &nodeCnt)
			if nodeCnt*(nodeCnt-1) == edgeCnt {
				ans++
			}
		}
	}
	return ans
}

// build graph
func buildGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return graph
}

// depth first search to count all connected component
func dfs(graph [][]int, currNode int, visited []bool, edgeCnt *int, nodeCnt *int) {
	*nodeCnt += 1
	*edgeCnt += len(graph[currNode])
	visited[currNode] = true
	for _, nei := range graph[currNode] {
		if !visited[nei] {
			dfs(graph, nei, visited, edgeCnt, nodeCnt)
		}
	}
}

func main() {
	println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
	println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}))
}
