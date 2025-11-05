package main

import "strings"

type TrieNode struct {
	IsEnd    bool
	Children map[string]*TrieNode
}

func removeSubfolders(folder []string) []string {
	root := &TrieNode{false, make(map[string]*TrieNode)}

	for _, path := range folder {
		curNode := root
		folderNames := strings.Split(path, "/")
		for _, folderName := range folderNames {
			if folderName == "" {
				continue
			}
			if _, ok := curNode.Children[folderName]; !ok {
				curNode.Children[folderName] = &TrieNode{false, make(map[string]*TrieNode)}
			}
			curNode = curNode.Children[folderName]
		}
		curNode.IsEnd = true
	}

	arr := []string{}

	for _, path := range folder {
		folderNames := strings.Split(path, "/")
		isSubFolder := false
		curNode := root
		for i, name := range folderNames {
			if name == "" {
				continue
			}
			nextNode := curNode.Children[name]
			if nextNode.IsEnd && i != len(folderNames)-1 {
				isSubFolder = true
				break
			}
			curNode = nextNode
		}
		if !isSubFolder {
			arr = append(arr, path)
		}
	}

	return arr
}

func main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	removeSubfolders(folder)
}
