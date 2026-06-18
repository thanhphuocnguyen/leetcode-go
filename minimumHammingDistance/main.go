package main

import "fmt"

type UnionFind struct {
	parent []int
	weight []int
}

func Constructor(n int) *UnionFind {
	parent := make([]int, n)
	weight := make([]int, n)
	for i := range n {
		parent[i] = i
		weight[i] = 1
	}

	return &UnionFind{parent, weight}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return
	}
	// if weight of px larger than weight py then assign it to the py to balance uf tree
	if uf.weight[px] > uf.weight[py] {
		// parent of px will be py
		uf.parent[px] = py
		uf.weight[py] += uf.weight[px]
	} else {
		uf.parent[py] = px
		uf.weight[px] += uf.weight[py]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	ans := 0
	uf := Constructor(n)
	for _, swap := range allowedSwaps {
		i, j := swap[0], swap[1]
		uf.union(i, j)
	}

	components := make(map[int]map[int]int, n)

	for i := range n {
		p := uf.find(i)
		if _, ok := components[p]; !ok {
			components[p] = map[int]int{}
		}
		components[p][source[i]]++
	}

	for i := range n {
		p := uf.find(i)
		freq := components[p]
		if val, ok := freq[target[i]]; ok && val > 0 {
			freq[target[i]]--
		} else {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumHammingDistance([]int{2, 3, 1}, []int{1, 2, 2}, [][]int{{0, 2}, {1, 2}}))
	fmt.Println(minimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}))
}
