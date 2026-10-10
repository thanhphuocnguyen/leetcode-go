package main

import "fmt"

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n-1)%2 != 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}
	maxBalance := m * n / 2
	memo := make([][][]int8, m)
	for i := range m {
		memo[i] = make([][]int8, n)
		for j := range n {
			memo[i][j] = make([]int8, maxBalance+1)
		}
	}

	var dfs func(r, c, balance int) bool
	dfs = func(r, c, balance int) bool {
		if grid[r][c] == '(' {
			balance += 1
		} else {
			balance -= 1
		}

		if balance < 0 || balance > maxBalance {
			return false
		}

		if memo[r][c][balance] != 0 {
			return memo[r][c][balance] == 2
		}

		if r == m-1 && c == n-1 {
			if balance == 0 {
				return true
			} else {
				return false
			}
		}

		ans := false
		if r+1 < m {
			ans = ans || dfs(r+1, c, balance)
		}

		if c+1 < n {
			ans = ans || dfs(r, c+1, balance)
		}
		if ans {
			memo[r][c][balance] = 2
		} else {
			memo[r][c][balance] = 1
		}

		return ans
	}
	return dfs(0, 0, 0)
}

func main() {
	fmt.Println(hasValidPath([][]byte{{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'}}))
}
