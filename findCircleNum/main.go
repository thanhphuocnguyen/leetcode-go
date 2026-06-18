package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)

	ans := 0

	for i := range n {
		if !visited[i] {
			ans++
			dfs(isConnected, visited, i)
		}
	}
	return ans
}

func dfs(isConnected [][]int, visited []bool, node int) {
	visited[node] = true

	for nei, status := range isConnected[node] {
		if !visited[nei] && status == 1 && isConnected[nei][node] == 1 {
			dfs(isConnected, visited, nei)
		}
	}
}

func main() {
	fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
	fmt.Println(findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
