package main

import (
	"fmt"
	"math"
)

type TrieNode struct {
	children [26]*TrieNode
	id       int
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: [26]*TrieNode{},
		id:       -1,
	}
}

func (tr *TrieNode) add(str string, id *int) int {
	node := tr
	for i := range str {
		b := int(str[i] - 'a')
		if node.children[b] == nil {
			node.children[b] = NewTrie()
		}
		node = node.children[b]
	}

	if node.id == -1 {
		node.id = *id
		*id++
	}
	return node.id
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n, m := len(source), len(original)

	// apply Floyd Warshall algo
	dist := make([][]int, 2*m)
	for i := range dist {
		dist[i] = make([]int, 2*m)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
		dist[i][i] = 0
	}
	id := 0
	// Build Trie
	root := NewTrie()
	for i := range m {
		u, v, c := original[i], changed[i], cost[i]
		// add each unique str with and ID
		uId := root.add(u, &id)
		vId := root.add(v, &id)
		// Init dist array for running Floyd Warshall algorithm
		dist[uId][vId] = min(dist[uId][vId], c)
	}

	// Run Floyd Warshall algorithm O(n^3)
	sz := id
	for k := range sz {
		for i := range sz {
			for j := range sz {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}

	// Apply dynamic programming or source length
	dp := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i < n+1; i++ {
		if dp[i-1] == math.MaxInt64 {
			continue
		}

		if target[i-1] == source[i-1] {
			dp[i] = min(dp[i], dp[i-1])
			continue
		}
		var u, v *TrieNode = root, root
		// run from i to n to find shortest path or each substring len, break when no found in Trie
		for j := i; j < n+1; j++ {
			u, v = u.children[source[j-1]-'a'], v.children[target[j-1]-'a']
			// if there are no string in side Trie break the loop
			if u == nil || v == nil {
				break
			}
			if u.id == -1 || v.id == -1 || dist[u.id][v.id] == math.MaxInt32 {
				continue
			}
			if dp[i-1] != math.MaxInt64 {
				dp[j] = min(dp[j], dp[i-1]+int64(dist[u.id][v.id]))
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(minimumCost("abcd", "acbe", []string{"a", "b", "c", "c", "e", "d"}, []string{"b", "c", "b", "e", "b", "e"}, []int{2, 5, 5, 1, 2, 20}))
	fmt.Println(minimumCost("abcdefgh", "acdeeghh", []string{"bcd", "fgh", "thh"}, []string{"cde", "thh", "ghh"}, []int{1, 3, 5}))
}
