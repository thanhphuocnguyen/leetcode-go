package main

import "fmt"

func main() {
	words := []string{"abc", "ab", "bc", "b"}
	fmt.Println(sumPrefixScores(words))
}

type trieNode struct {
	cnt      int
	children [26]*trieNode
}

func sumPrefixScores(words []string) []int {
	trie := newTrieNode()
	trie.insert(words)
	sps := make([]int, len(words))
	for i, w := range words {
		sps[i] = trie.sumPrefixCnt(w)
	}
	return sps
}

func newTrieNode() *trieNode {
	return &trieNode{
		cnt:      1,
		children: [26]*trieNode{},
	}
}

func (t *trieNode) insert(words []string) {
	for _, w := range words {
		node := t
		for _, c := range w {
			cIdx := c - 'a'
			if node.children[cIdx] == nil {
				node.children[cIdx] = newTrieNode()
			} else {
				node.children[cIdx].cnt++
			}
			node = node.children[cIdx]
		}
	}
}

func (t *trieNode) sumPrefixCnt(word string) int {
	node := t
	ans := 0
	for _, c := range word {
		cIdx := c - 'a'
		if node.children[cIdx] == nil {
			break
		}
		ans += node.children[cIdx].cnt
		node = node.children[cIdx]
	}
	return ans
}
