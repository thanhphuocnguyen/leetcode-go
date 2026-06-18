package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	seen := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if seen[node] != nil {
			return seen[node]
		}
		cloned := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
		seen[node] = cloned
		for i, nei := range node.Neighbors {
			cloned.Neighbors[i] = dfs(nei)
		}
		return cloned
	}
	return dfs(node)
}

func main() {
	// [[2,4],[1,3],[2,4],[1,3]]
	graph := &Node{Val: 1, Neighbors: []*Node{
		{Val: 2, Neighbors: []*Node{
			{Val: 1},
			{Val: 3},
		}},
		{Val: 4, Neighbors: []*Node{
			{Val: 1},
			{Val: 3},
		}},
	}}
	clonedGraph := cloneGraph(graph)
	println(clonedGraph.Val) // Output: 1
}
