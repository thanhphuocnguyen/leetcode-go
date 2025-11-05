package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
}

type Trie struct {
	children [10]*Trie
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	trie := newTrie()
	for _, num := range arr1 {
		trie.insert(num)
	}
	lp := 0
	for _, num := range arr2 {
		lp = max(lp, trie.findLongestPrefix(num))
	}
	return lp
}

func newTrie() *Trie {
	return &Trie{children: [10]*Trie{}}
}

func (t *Trie) insert(num int) {
	node := t
	numStr := strconv.Itoa(num)
	for _, ch := range numStr {
		digit := ch - '0'
		if node.children[digit] == nil {
			node.children[digit] = newTrie()
		}
		node = node.children[digit]
	}
}

func (t *Trie) findLongestPrefix(num int) int {
	node := t
	numStr := strconv.Itoa(num)
	len := 0
	for _, ch := range numStr {
		digit := ch - '0'
		if node.children[digit] == nil {
			break
		}
		node = node.children[digit]
		len++
	}
	return len
}
