package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
	root := this
	for _, r := range word {
		ord := int(r - 'a')
		if root.children[ord] == nil {
			return false
		}
		root = root.children[ord]
	}
	return root.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, r := range prefix {
		ord := int(r - '0')
		if root.children[ord] == nil {
			return false
		}
		root = root.children[ord]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
