package main

import (
	"fmt"
	"sort"
)

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
	if uf.weight[px] > uf.weight[py] {
		uf.parent[px] = py
		uf.weight[py] += uf.weight[px]
	} else {
		uf.parent[py] = px
		uf.weight[px] += uf.weight[py]
	}
}

type RuneMapList struct {
	runes []byte
	idx   int
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := Constructor(n)
	for _, p := range pairs {
		uf.union(p[0], p[1])
	}

	// take components
	components := make(map[int]*RuneMapList)
	for i := range n {
		f := uf.find(i)
		if _, ok := components[f]; !ok {
			components[f] = &RuneMapList{runes: []byte{}, idx: 0}
		}
		components[f].runes = append(components[f].runes, s[i])
	}
	for _, comp := range components {
		sort.Slice(comp.runes, func(i, j int) bool {
			return comp.runes[i] < comp.runes[j]
		})
	}
	ans := make([]byte, len(s))
	// map to ans
	for i := range n {
		p := uf.find(i)
		ans[i] = components[p].runes[components[p].idx]
		components[p].idx++
	}
	return string(ans)
}

func main() {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
}
