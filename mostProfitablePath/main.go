package main

import (
	"fmt"
	"math"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	adj := make([][]int, len(amount))
	for i := range len(amount) {
		adj[i] = make([]int, 0)
	}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	bobTimestamps := make([]int, len(amount))
	for i := range len(amount) {
		bobTimestamps[i] = -1
	}
	bobPath := make([]int, 0)
	var getBobPath func(node, parentNode int) bool
	getBobPath = func(node, parentNode int) bool {
		if node == 0 {
			return true
		}
		for _, nei := range adj[node] {
			if nei != parentNode {
				bobPath = append(bobPath, node)
				if getBobPath(nei, node) {
					return true
				}
				bobPath = bobPath[:len(bobPath)-1]
			}
		}
		return false
	}
	getBobPath(bob, -1)
	for i, node := range bobPath {
		bobTimestamps[node] = i
	}
	return dfs(adj, bobTimestamps, amount, 0, -1, 0, 0)
}

func dfs(adj [][]int, bobTimestamps, amount []int, node, parentNode, currentTimestamp, currentProfit int) int {
	nodeProfit := amount[node]
	if bobTimestamps[node] == currentTimestamp {
		nodeProfit /= 2
	} else if bobTimestamps[node] != -1 && currentTimestamp > bobTimestamps[node] {
		nodeProfit = 0
	}
	currentProfit += nodeProfit
	if len(adj[node]) == 1 && node != 0 {
		return currentProfit
	}
	maxProfit := math.MinInt
	for _, nei := range adj[node] {
		if nei != parentNode {
			maxProfit = max(maxProfit, dfs(adj, bobTimestamps, amount, nei, node, currentTimestamp+1, currentProfit))
		}
	}
	return maxProfit
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
	bob := 3
	amount := []int{-2, 4, 2, -4, 6}
	fmt.Print(mostProfitablePath(edges, bob, amount))
}
