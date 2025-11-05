package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adjList := make([][]int, n)
	for i := range n {
		adjList[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}
	componentCount := 0
	dfs(0, -1, adjList, values, k, &componentCount)
	return componentCount
}

func dfs(currNode, parentNode int, adjList [][]int, values []int, k int, componentCount *int) int {
	sum := 0
	for _, neighbor := range adjList[currNode] {
		if neighbor != parentNode {
			sum += dfs(neighbor, currNode, adjList, values, k, componentCount)
			sum %= k
		}
	}
	sum += values[currNode]
	sum %= k
	if sum == 0 {
		*componentCount++
	}
	return sum
}

func main() {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}}
	values := []int{1, 2, 3, 4, 5, 6}
	k := 3
	maxKDivisibleComponents(n, edges, values, k)
}
