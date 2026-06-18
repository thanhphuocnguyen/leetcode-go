package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (this *Trie) add(word string) {
	root := this
	for _, r := range word {
		ord := int(r - 'a')
		if root.children[ord] == nil {
			root.children[ord] = &Trie{
				children: [26]*Trie{},
				isEnd:    false,
			}
		}
		root = root.children[ord]
	}
	root.isEnd = true
}

func (this *Trie) findWord(word string) bool {
	root := this
	n := len(word)
	var dfs func(root *Trie, idx int) bool
	dfs = func(root *Trie, idx int) bool {
		if root == nil {
			return false
		}
		if idx == n {
			return root.isEnd
		}

		if word[idx] == '.' {
			for i := range 26 {
				node := root.children[i]
				if node != nil && dfs(node, idx+1) {
					return true
				}
			}
		} else {
			ord := int(word[idx] - 'a')
			node := root.children[ord]
			if dfs(node, idx+1) {
				return true
			}
		}

		return false
	}

	return dfs(root, 0)
}

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	trie := &Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
	return WordDictionary{trie}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.add(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.findWord(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	myWS := Constructor()
	myWS.AddWord("bad")
	myWS.AddWord("dad")
	myWS.AddWord("mad")
	fmt.Println(myWS.Search("pad"))
	fmt.Println(myWS.Search("bad"))
	fmt.Println(myWS.Search(".ad"))
	fmt.Println(myWS.Search("b.."))
}
