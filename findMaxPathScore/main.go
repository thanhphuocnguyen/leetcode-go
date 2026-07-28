package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
)

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	G := make([][][2]int, n)
	for i := range n {
		G[i] = make([][2]int, 0)
	}
	mp := make(map[int]bool)
	for _, e := range edges {
		u, v, c := e[0], e[1], e[2]
		G[u] = append(G[u], [2]int{v, c})
		if !mp[c] {
			mp[c] = true
		}
	}
	weights := make([]int, 0, len(mp))
	for key := range mp {
		weights = append(weights, key)
	}
	sort.Ints(weights)
	left, right := 0, len(weights)-1
	order := topologicalSort(G, n)
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		threshold := weights[mid]
		if feasible(G, online, order, n, threshold, k) {
			ans = threshold
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return ans
}

func feasible(graph [][][2]int, online []bool, order []int, n, threshold int, k int64) bool {
	dist := make([]int64, n)
	for i := range n {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	for _, node := range order {
		if dist[node] == math.MaxInt64 {
			continue
		}
		for _, edge := range graph[node] {
			v, c := edge[0], edge[1]
			if c < threshold {
				continue
			}

			if v != n-1 && !online[v] {
				continue
			}

			if dist[node]+int64(c) < dist[v] {
				dist[v] = dist[node] + int64(c)
			}
		}
	}
	return dist[n-1] <= k
}

func topologicalSort(graph [][][2]int, n int) []int {
	indegree := make([]int, n)
	for i := range n {
		for _, v := range graph[i] {
			indegree[v[0]]++
		}
	}
	q := list.New()
	for u, deg := range indegree {
		if deg == 0 {
			q.PushBack(u)
		}
	}

	ans := make([]int, 0)
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		ans = append(ans, u)
		for _, v := range graph[u] {
			indegree[v[0]]--
			if indegree[v[0]] == 0 {
				q.PushBack(v[0])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(findMaxPathScore([][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}}, []bool{true, true, true, true}, 10))
}
