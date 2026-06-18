package main

import (
	"fmt"
)

func minimumDistance(word string) int {
	n := len(word)
	memo := make([][][]int, 27)
	for i := range 27 {
		memo[i] = make([][]int, 27)
		for j := range 27 {
			memo[i][j] = make([]int, n+1)
			for k := range n + 1 {
				memo[i][j][k] = -1
			}
		}
	}
	return dp(memo, word, 26, 26, 0)
}

// will be all by left, all by right, and left right s
func dp(memo [][][]int, word string, left, right, pos int) int {
	if pos == len(word) {
		return 0
	}

	if memo[left][right][pos] != -1 {
		return memo[left][right][pos]
	}
	curr := int(word[pos] - 'A')
	// type curr char by left and -> then switch to right hand
	leftFinger := dist(left, curr) + dp(memo, word, curr, right, pos+1)
	// type curr char by right and switch to left hand
	rightFinger := dist(right, curr) + dp(memo, word, left, curr, pos+1)
	memo[left][right][pos] = min(leftFinger, rightFinger) // type curr char by right finger

	return memo[left][right][pos]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(prev, cur int) int {
	if prev == 26 || cur == 26 {
		return 0
	}

	return abs(prev/6-cur/6) + abs(prev%6-cur%6)
}

func main() {
	fmt.Println(minimumDistance("HAPPY"))
	fmt.Println(minimumDistance("CAKE"))
}
