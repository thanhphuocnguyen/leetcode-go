package main

import "fmt"

func constructProductMatrix(grid [][]int) [][]int {
	const MOD = 12345
	n, m := len(grid), len(grid[0])
	flatten := make([]int, n*m)
	for i, row := range grid {
		for j, cell := range row {
			flatten[i*m+j] = cell
		}
	}
	prefixProd := make([]int, n*m)
	suffixProd := make([]int, n*m)
	prefixProd[0] = 1
	suffixProd[n*m-1] = 1
	for i := 1; i < n*m; i++ {
		prefixProd[i] = (prefixProd[i-1] * flatten[i-1]) % MOD
	}
	for i := n*m - 2; i >= 0; i-- {
		suffixProd[i] = (suffixProd[i+1] * flatten[i+1]) % MOD
	}
	ans := make([][]int, n)
	for i := range n {
		ans[i] = make([]int, m)
		for j := range m {
			ans[i][j] = (prefixProd[i*m+j] * suffixProd[i*m+j]) % MOD
		}
	}

	return ans
}

func main() {
	fmt.Println(constructProductMatrix([][]int{{10, 20}, {18, 16}, {17, 14}, {16, 9}, {14, 6}, {16, 5}, {14, 8}, {20, 13}, {16, 10}, {14, 17}}))
	fmt.Println(constructProductMatrix([][]int{{1, 2, 2}, {1, 4, 3}}))
	fmt.Println(constructProductMatrix([][]int{{1, 2, 2}, {1, 4, 3}}))
	fmt.Println(constructProductMatrix([][]int{{12345}, {2}, {1}}))
	fmt.Println(constructProductMatrix([][]int{{1, 2}, {3, 4}}))
}
