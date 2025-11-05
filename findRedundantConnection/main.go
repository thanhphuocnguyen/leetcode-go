package main

import "fmt"

type DisjoinSet struct {
	Parent []int
	Size   []int
}

func initDSU(n int) *DisjoinSet {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	return &DisjoinSet{
		Parent: parent,
		Size:   size,
	}
}

func (dsu *DisjoinSet) Union(u, v int) bool {
	uParent := dsu.Find(u)
	vParent := dsu.Find(v)
	if uParent == vParent {
		return false
	}
	uSize := dsu.Size[u]
	vSize := dsu.Size[v]
	if uSize < vSize {
		dsu.Parent[uParent] = vParent
		dsu.Size[vParent] += uSize
	} else {
		dsu.Parent[vParent] = uParent
		dsu.Size[uParent] += vSize
	}
	return true
}

func (dsu *DisjoinSet) Find(i int) int {
	if dsu.Parent[i] != i {
		dsu.Parent[i] = dsu.Find(dsu.Parent[i])
	}
	return dsu.Parent[i]
}

func findRedundantConnection(edges [][]int) []int {
	dsu := initDSU(len(edges))
	for _, edge := range edges {
		if !dsu.Union(edge[0]-1, edge[1]-1) {
			return []int{edge[0], edge[1]}
		}
	}
	return []int{}
}

func main() {
	edges := [][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}}
	fmt.Println(findRedundantConnection(edges))
}
