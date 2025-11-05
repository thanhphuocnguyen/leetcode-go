package main

import "sort"

type NumberContainers struct {
	indexNum   map[int]int
	numIndices map[int][]int
	sorted     map[int]bool
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexNum:   make(map[int]int),
		numIndices: make(map[int][]int),
		sorted:     make(map[int]bool),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if val, ok := this.indexNum[index]; ok {
		for i := 0; i < len(this.numIndices[val]); i++ {
			if this.numIndices[val][i] == index {
				this.numIndices[val] = append(this.numIndices[val][:i], this.numIndices[val][i+1:]...)
				break
			}
		}
	}
	this.indexNum[index] = number
	if _, ok := this.numIndices[number]; !ok {
		this.numIndices[number] = []int{}
	}
	this.numIndices[number] = append(this.numIndices[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if _, ok := this.numIndices[number]; !ok {
		return -1
	}
	if _, ok := this.sorted[number]; !ok {
		this.sorted[number] = true
		sort.Ints(this.numIndices[number])
	}
	return this.numIndices[number][0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
