package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

var ans = 0

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	dfs(root, 1)
	return ans
}

func dfs(root *Node, depth int) {
	if len(root.Children) == 0 {
		ans = max(ans, depth)
		return
	}
	for _, node := range root.Children {
		dfs(node, depth+1)
	}
}

func main() {
	// Example usage
	root := &Node{Val: 1, Children: []*Node{
		{Val: 2, Children: nil},
		{Val: 3, Children: []*Node{
			{Val: 6, Children: nil},
			{Val: 7, Children: nil},
		}},
		{Val: 4, Children: nil},
		{Val: 5, Children: nil},
	}}
	println(maxDepth(root)) // Output should be 3
}
