package main

import (
	"fmt"
	"math"
)

type Trie struct {
	children [26]*Trie
	isWord   bool
	idx      int
}

func (this *Trie) AddReverse(word string, idx int) {
	root := this
	// add in reversed to make the problem find suffix to find prefix
	for i := len(word) - 1; i >= 0; i-- {
		bIdx := word[i] - 'a'

		if root.children[bIdx] == nil {
			root.children[bIdx] = &Trie{children: [26]*Trie{}, idx: -1}
		}
		root = root.children[bIdx]
	}

	root.isWord = true
	if root.idx == -1 {
		root.idx = idx
	}

	// cache the shortest word in the container for edge case (not found any prefix)

}

func (this *Trie) PrefixFind(suffix string) int {
	root := this
	// travel in reversed
	for i := len(suffix) - 1; i >= 0; i-- {
		bIdx := suffix[i] - 'a'
		if root.children[bIdx] == nil {
			break
		}
		root = root.children[bIdx]
	}

	// if current prefix isWord just return its idx
	if root.isWord {
		return root.idx
	} else {
		// bfs
		q := make([]*Trie, 0)
		q = append(q, root)
		found := make([]int, 0)
		for len(q) > 0 {
			sz := len(q)
			q2 := make([]*Trie, 0)
			for i := range sz {
				node := q[i]
				for _, child := range node.children {
					// skip nil node
					if child == nil {
						continue
					}

					// find word in current level
					if child.isWord {
						found = append(found, child.idx)
					} else {
						q2 = append(q2, child)
					}
				}
				// skip next level if we found current level
			}
			if len(found) > 0 {
				q = q[:0]
				break
			}
			q = q2
		}

		// find min idx for longest common suffix
		minIdx := math.MaxInt32
		for _, idx := range found {
			minIdx = min(idx, minIdx)
		}

		return minIdx
	}
}

func Constructor() Trie {
	return Trie{
		children: [26]*Trie{},
		idx:      -1,
	}
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	trie := Constructor()
	for idx, word := range wordsContainer {
		trie.AddReverse(word, idx)
	}
	cache := map[string]int{}
	ans := make([]int, len(wordsQuery))
	for idx, q := range wordsQuery {
		if _, ok := cache[q]; ok {
			ans[idx] = cache[q]
		} else {
			ans[idx] = trie.PrefixFind(q)
		}
	}

	return ans
}

func main() {
	fmt.Println(stringIndices([]string{"abc", "bca", "cab", "abb", "bac"}, []string{"a", "b", "c", "aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc", "aaa", "aab", "aac", "aba", "abb", "abc", "aca", "acb", "acc", "baa", "bab", "bac", "bba", "bbb", "bbc", "bca", "bcb", "bcc", "caa", "cab", "cac", "cba", "cbb", "cbc", "cca", "ccb", "ccc"}))
	fmt.Println(stringIndices([]string{"abcde", "abcde"}, []string{"abcde", "bcde", "cde", "de", "e"}))
	fmt.Println(stringIndices([]string{"abcdefgh", "poiuygh", "ghghgh"}, []string{"gh", "acbfgh", "acbfegh"}))
	fmt.Println(stringIndices([]string{"abcd", "bcd", "xbcd"}, []string{"cd", "bcd", "xyz"}))
}
