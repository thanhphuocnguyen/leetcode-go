package main

type Node struct {
	Key    int
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
	Color  Color // RED for red, BLACK for black
}
type Color int

const (
	RED Color = iota
	BLACK
)

type RedBlackTree struct {
	Root *Node
	Nil  *Node
}

func Constructor() RedBlackTree {
	nilNode := &Node{Color: BLACK} // Sentinel node
	nilNode.Left = nilNode
	nilNode.Right = nilNode
	return RedBlackTree{
		Nil:  nilNode,
		Root: nilNode,
	}
}

func (tree *RedBlackTree) rotateLeft(node *Node) {
	rotateNode := node.Right     // Node to rotate up
	node.Right = rotateNode.Left // Move left subtree of rotateNode to node's right
	if rotateNode.Left != tree.Nil {
		// Update parent of the left subtree of rotateNode
		rotateNode.Left.Parent = node
	}
	rotateNode.Parent = node.Parent // Update parent of rotateNode
	if node.Parent == tree.Nil {
		// If node is root, update tree root
		tree.Root = rotateNode
	} else if node == node.Parent.Left {
		// If node is a left child, update parent's left pointer
		node.Parent.Left = rotateNode
	} else {
		// If node is a right child, update parent's right pointer
		node.Parent.Right = rotateNode
	}
	// Move node down and make it left child of rotateNode
	rotateNode.Left = node
	// Update parent of node
	node.Parent = rotateNode
}

/*
rotate right summary
  - we want to rotate node to the right
  - so the left node of current not gonna be parent of the node
  - and the right subtree of the left node will be appended to the left of the current node
  - the left node's parent now is the current node's parent
  - if the tree root is nil so the left node now will be root
  - else if the current node is the current node parent left -> assign left node to it
  - do for case current node is the current node parent right
  - left node right will be current node
  - current node parent will be left node
*/
func (tree *RedBlackTree) rotateRight(node *Node) {
	rotateNode := node.Left      // Node to rotate up
	node.Left = rotateNode.Right // Move right subtree of rotateNode to node's left
	if rotateNode.Right != tree.Nil {
		// Update parent of the right subtree of rotateNode
		rotateNode.Right.Parent = node
	}
	rotateNode.Parent = node.Parent // Update parent of rotateNode
	if node.Parent == tree.Nil {
		// If node is root, update tree root
		tree.Root = rotateNode
	} else if node == node.Parent.Left {
		// If node is a left child, update parent's left pointer
		node.Parent.Left = rotateNode
	} else {
		// If node is a right child, update parent's right pointer
		node.Parent.Right = rotateNode
	}
	// Move node down and make it left child of rotateNode
	rotateNode.Right = node
	// Update parent of node
	node.Parent = rotateNode
}

func (tree *RedBlackTree) fixupTree(node *Node) {
	for node.Parent.Color == RED {
		// handle for left tree
		parent := node.Parent
		grandParent := parent.Parent
		if parent == grandParent.Left {
			// Get uncle node
			uncle := grandParent.Right
			if uncle.Color == RED {
				// Case 1: Uncle is red, recolor
				parent.Color = BLACK
				uncle.Color = BLACK
				grandParent.Color = RED
				// Move up the tree
				node = grandParent
			} else {
				if node == parent.Right {
					// Case 2: Node is right child, rotate left
					node = parent
					tree.rotateLeft(node)
				}
				// Update parent and grandparent after rotation
				parent.Color = BLACK
				grandParent.Color = RED
				tree.rotateRight(grandParent)
			}
		} else {
			// parent is on the right subtree
			uncle := grandParent.Left
			if uncle.Color == BLACK {
				// case 2: uncle is red
				// recolor grandparent, parent and uncle
				grandParent.Color = RED
				parent.Color = BLACK
				uncle.Color = BLACK
				// move up to grandparent
				node = grandParent
			} else {
				// check for case 3:
				// case uncle is black but the node parent and grand parent form a triangle
				if parent.Left == node {
					// we rotate right parent to form case 4
					node = parent
					tree.rotateRight(node)
				}
				// case 4: the grandparent, parent and uncle form a line and uncle is black
				// this case we recolor grandparent to red, and recolor parent to black and rotate left the grandparent
				parent.Color = BLACK
				grandParent.Color = RED
				tree.rotateLeft(grandParent)
			}
		}
	}
	// the root is always black
	tree.Root.Color = BLACK
}

// findNode is a helper function to find a node with the given key in the tree.
func (tree *RedBlackTree) findNode(key int) *Node {
	// This function is used to find a node with the given key in the tree.
	node := tree.Root
	// Traverse the tree until we find the node or reach the sentinel node (Nil).
	for node != tree.Nil {
		// Compare the key with the current node's key to decide whether to go left, right, or if we found the node.
		if key < node.Key {
			// If the key is less than the current node's key, go to the left subtree.
			node = node.Left
		} else if key > node.Key {
			// If the key is greater than the current node's key, go to the right subtree.
			node = node.Right
		} else {
			// If the key matches the current node's key, we have found the node.
			return node
		}
	}
	return tree.Nil
}

func transplant(tree *RedBlackTree, u, v *Node) {
	// Transplant node v in place of node u in the tree.
	if u.Parent == tree.Nil {
		// If u is the root, update the tree's root to be v.
		tree.Root = v
	} else if u == u.Parent.Left {
		// If u is a left child, update the parent's left pointer to point to v.
		u.Parent.Left = v
	} else {
		// If u is a right child, update the parent's right pointer to point to v.
		u.Parent.Right = v
	}
	// Update the parent of node v to be the parent of node u.
	v.Parent = u.Parent
}

func findMinimum(tree *RedBlackTree, node *Node) *Node {
	// Find the node with the minimum key in the subtree rooted at the given node.
	current := node
	// Traverse the left children until we reach the sentinel node (Nil).
	for current.Left != tree.Nil {
		current = current.Left
	}
	return current
}

func (tree *RedBlackTree) fixupDelete(node *Node) {
	for node != tree.Root && node.Color == BLACK {
		if node == node.Parent.Left {
			sibling := node.Parent.Right
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.Parent.Color = RED
				tree.rotateLeft(node.Parent)
				sibling = node.Parent.Right
			}
			if sibling.Left.Color == BLACK && sibling.Right.Color == BLACK {
				sibling.Color = RED
				node = node.Parent
			} else {
				if sibling.Right.Color == BLACK {
					sibling.Left.Color = BLACK
					sibling.Color = RED
					tree.rotateRight(sibling)
					sibling = node.Parent.Right
				}
				sibling.Color = node.Parent.Color
				node.Parent.Color = BLACK
				sibling.Right.Color = BLACK
				tree.rotateLeft(node.Parent)
				node = tree.Root
			}
		} else {
			sibling := node.Parent.Left
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.Parent.Color = RED
				tree.rotateRight(node.Parent)
				sibling = node.Parent.Left
			}
			if sibling.Left.Color == BLACK && sibling.Right.Color == BLACK {
				sibling.Color = RED
				node = node.Parent
			} else {
				if sibling.Left.Color == BLACK {
					sibling.Right.Color = BLACK
					sibling.Color = RED
					tree.rotateLeft(sibling)
					sibling = node.Parent.Left
				}
				sibling.Color = node.Parent.Color
				node.Parent.Color = BLACK
				sibling.Left.Color = BLACK
				tree.rotateRight(node.Parent)
				node = tree.Root
			}
		}
	}
	node.Color = BLACK
}

func (tree *RedBlackTree) Insert(val int) {
	// newNode := &Node{Value: val, Left: }
	newNode := &Node{Value: val, Color: RED, Left: tree.Nil, Right: tree.Nil}
	curr := tree.Root
	var parent *Node
	for curr != tree.Nil {
		parent = curr
		if newNode.Value < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	newNode.Parent = parent
	if parent == tree.Nil {
		tree.Root = newNode
	} else if newNode.Value < parent.Value {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	tree.fixupTree(newNode)
}

func (tree *RedBlackTree) Get(key int) int {
	node := tree.Root
	for node != tree.Nil {
		if key < node.Key {
			node = node.Left
		} else if key > node.Key {
			node = node.Right
		} else {
			return node.Value
		}
	}
	return -1
}

func (tree *RedBlackTree) Delete(key int) {
	node := tree.findNode(key)
	if node == tree.Nil {
		// Node with the given key does not exist, nothing to delete.
		return
	}
	var subtree *Node
	var originalColor Color
	// case 1: node left is nil
	if node.Left == tree.Nil {
		subtree = node.Right
		transplant(tree, node, node.Right)
	} else if node.Right == tree.Nil {
		// case 2: node right is nil
		subtree = node.Left
		transplant(tree, node, node.Left)
	} else {
		// case 3: node has two children
		// find the minimum node in the right subtree
		minNode := findMinimum(tree, node.Right)
		// save the original color of the minimum node
		originalColor = minNode.Color
		subtree = minNode.Right
		if minNode.Parent == node {
			// If the minimum node is a direct child of the node to be deleted,
			// we can simply transplant the minimum node's right child to the minimum node's position.
			subtree.Parent = minNode
		} else {
			// If the minimum node is not a direct child,
			// we need to transplant the minimum node's right child to the minimum node's position,
			// and then transplant the minimum node to the node's position.
			transplant(tree, minNode, minNode.Right)
			minNode.Right = node.Right
			minNode.Right.Parent = minNode
		}
		// Transplant the minimum node to the position of the node to be deleted.
		transplant(tree, node, minNode)
		minNode.Left = node.Left
		minNode.Left.Parent = minNode
		minNode.Color = node.Color
	}
	if originalColor == BLACK {
		tree.fixupDelete(subtree)
	}
}
