package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	r      int
	c      int
	mask   int
	energy int
	steps  int
}

var DIRECTIONS = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func minMoves(classroom []string, energy int) int {
	m, n := len(classroom), len(classroom[0])
	lPos := make([][2]int, 0)
	sr, sc := -1, -1
	cnt := 0
	lIndexes := make([][]int, m)
	for i, row := range classroom {
		lIndexes[i] = make([]int, n)
		for j, cell := range row {
			switch cell {
			case 'S':
				sr = i
				sc = j
			case 'L':
				lIndexes[i][j] = cnt
				cnt++
				lPos = append(lPos, [2]int{i, j})
			}
		}
	}

	q := list.New()
	q.PushBack(&Node{
		r:      sr,
		c:      sc,
		mask:   0,
		energy: energy,
		steps:  0,
	})
	allMask := (1 << cnt) - 1
	bestEnergy := make([][][]int, m)
	for i := range m {
		bestEnergy[i] = make([][]int, n)
		for j := range n {
			bestEnergy[i][j] = make([]int, allMask+1)
			for k := range allMask + 1 {
				bestEnergy[i][j][k] = -1
			}
		}
	}
	bestEnergy[sr][sc][0] = energy
	for q.Len() > 0 {
		node := q.Remove(q.Front())
		val := node.(*Node)
		r, c, mask, currEnergy, steps := val.r, val.c, val.mask, val.energy, val.steps
		if mask == allMask {
			return val.steps
		}

		for _, d := range DIRECTIONS {
			nr, nc := r+d[0], c+d[1]
			if isValid(nr, nc, m, n) && classroom[nr][nc] != 'X' {
				newEnergy := currEnergy - 1
				if newEnergy < 0 {
					continue
				}
				newMask := mask

				switch classroom[nr][nc] {
				case 'L':
					newMask = newMask | (1 << lIndexes[nr][nc])
				case 'R':
					newEnergy = energy
				}

				if bestEnergy[nr][nc][newMask] >= newEnergy {
					continue
				}

				bestEnergy[nr][nc][newMask] = newEnergy

				q.PushBack(&Node{
					r:      nr,
					c:      nc,
					mask:   newMask,
					energy: newEnergy,
					steps:  steps + 1,
				})
			}
		}
	}
	return -1
}

func isValid(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func main() {
	fmt.Println(minMoves([]string{"LS", "RL"}, 4))
}
