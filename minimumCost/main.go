package main

type DisjointSet struct {
	Parents []int
	Weights []int
}

func DisjointSetConstructor(n int) *DisjointSet {
	parents := make([]int, n)
	weights := make([]int, n)
	for i := range n {
		parents[i] = i + 1
		weights[i] = 1
	}
	return &DisjointSet{
		Parents: parents,
		Weights: make([]int, 0),
	}
}

func (ds *DisjointSet) Find(x int) int {
	root := ds.Parents[x]
	if root != ds.Parents[root] {
		ds.Parents[x] = ds.Find(root)
		return ds.Parents[x]
	}
	return root
}

func (ds *DisjointSet) Union(x, y int) {
	xRoot, yRoot := ds.Find(x), ds.Find(y)
	xWeight, yWeight := ds.Weights[x], ds.Weights[y]
	if xRoot == yRoot {
		return
	}
	if xWeight < yWeight {
		ds.Parents[xRoot] = yRoot
		ds.Weights[yRoot] += ds.Weights[xRoot]
	} else {
		ds.Parents[yRoot] = xRoot
		ds.Weights[xRoot] += ds.Weights[yRoot]
	}
}
func minimumCost(n int, edges [][]int, query [][]int) []int {

	minCosts := make([]int, n)
	for i := range minCosts {
		minCosts[i] = -1
	}
	disjointSet := DisjointSetConstructor(n)
	for _, edge := range edges {
		disjointSet.Union(edge[0], edge[1])
	}

	for _, edge := range edges {
		root := disjointSet.Find(edge[0])
		if minCosts[root] == -1 {
			minCosts[root] = edge[2]
		} else {
			minCosts[root] &= edge[2]
		}
	}
	ans := make([]int, 0)
	for _, q := range query {
		start, end := disjointSet.Find(q[0]), disjointSet.Find(q[1])
		if start != end {
			ans = append(ans, -1)
		} else {
			ans = append(ans, minCosts[start])
		}
	}
	return ans
}
