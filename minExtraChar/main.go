package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
}

type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

func minExtraChar(s string, dictionary []string) int {
	root := buildTrie(dictionary)
	n := len(s)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	return dp(0, n, s, memo, root)
}

func dp(start int, n int, s string, memo []int, trie *TrieNode) int {
	if start == n {
		return 0
	}

	if memo[start] != -1 {
		return memo[start]
	}
	ans := dp(start+1, n, s, memo, trie) + 1
	node := trie
	for end := start; end < n; end++ {
		c := rune(s[end])
		if _, ok := node.Children[c]; !ok {
			break
		}
		node = node.Children[c]
		if node.IsWord {
			dpEnd := dp(end+1, n, s, memo, trie)
			if dpEnd != math.MaxInt32 {
				ans = min(ans, dpEnd)
			}
		}
	}
	memo[start] = ans
	return memo[start]
}

func buildTrie(dictionary []string) *TrieNode {
	root := &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsWord:   false,
	}
	for _, word := range dictionary {
		node := root
		for _, c := range word {
			if _, ok := node.Children[c]; !ok {

				node.Children[c] = &TrieNode{
					Children: make(map[rune]*TrieNode),
					IsWord:   false,
				}
			}
			node = node.Children[c]
		}
		node.IsWord = true
	}

	return root
}
