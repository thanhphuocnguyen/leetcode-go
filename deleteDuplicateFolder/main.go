package main

import (
	"fmt"
	"sort"
	"strings"
)

type Trie struct {
	serial   string
	children map[string]*Trie
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := &Trie{children: make(map[string]*Trie)}
	for _, path := range paths {
		curr := root
		for _, folder := range path {
			if _, ok := curr.children[folder]; !ok {
				curr.children[folder] = &Trie{
					children: map[string]*Trie{},
				}
			}
			curr = curr.children[folder]
		}
	}
	freq := make(map[string]int)
	var construct func(root *Trie)
	construct = func(root *Trie) {
		if len(root.children) == 0 {
			return
		}
		v := make([]string, 0, len(root.children))
		for folder, child := range root.children {
			construct(child)
			str := folder
			if len(child.serial) > 0 {
				str += fmt.Sprintf("(%s)", child.serial)
			}
			v = append(v, str)
		}
		sort.Strings(v)
		serial := strings.Join(v, "")
		root.serial = serial
		freq[serial]++
	}

	construct(root)
	ans := make([][]string, 0)
	path := make([]string, 0)
	var operate func(root *Trie)
	operate = func(root *Trie) {
		if freq[root.serial] > 1 {
			return
		}
		if len(path) > 0 {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		for folder, child := range root.children {
			path = append(path, folder)
			operate(child)
			path = path[:len(path)-1]
		}
	}
	operate(root)
	return ans
}

func main() {
	fmt.Println(deleteDuplicateFolder([][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}}))
}
