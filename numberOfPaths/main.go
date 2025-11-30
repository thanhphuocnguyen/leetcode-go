package main

import "fmt"

const MOD = 1_000_000_007

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range m {
		memo[i] = make([][]int, n)
		for j := range n {
			memo[i][j] = make([]int, k)
			for l := range k {
				memo[i][j][l] = -1
			}
		}
	}

	return dp(memo, grid, k, 0, 0, 0, m, n)
}

func dp(memo [][][]int, grid [][]int, k, mod, i, j, m, n int) int {
	if i >= m || j >= n {
		return 0
	}

	if memo[i][j][mod] != -1 {
		return memo[i][j][mod]
	}

	if i == m-1 && j == n-1 && (mod+grid[i][j])%k == 0 {
		return 1
	}

	cnt := 0
	newMod := (mod + grid[i][j]) % k
	cnt += dp(memo, grid, k, newMod, i+1, j, m, n)
	cnt += dp(memo, grid, k, newMod, i, j+1, m, n)

	memo[i][j][mod] = cnt % MOD
	return memo[i][j][mod]
}

func main() {
	fmt.Println(numberOfPaths([][]int{{1}, {5}, {3}, {7}, {3}, {2}, {3}, {5}}, 29))
	fmt.Println(numberOfPaths([][]int{{1, 5, 3, 7, 3, 2, 3, 5}}, 29))
	fmt.Println(numberOfPaths([][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3))
}
