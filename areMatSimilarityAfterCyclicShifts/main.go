package main

import "fmt"

func areSimilar(mat [][]int, k int) bool {
	m, n := len(mat), len(mat[0])
	k = k % n
	shifted := make([][]int, m)
	for i := range m {
		shifted[i] = make([]int, n)
		for j := range n {
			if i%2 == 0 {
				shifted[i][(j-k+n)%n] = mat[i][j]
			} else {
				shifted[i][(j+k)%n] = mat[i][j]
			}
		}
		for j := range n {
			if shifted[i][j] != mat[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(areSimilar([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2))
	fmt.Println(areSimilar([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 4))
}
