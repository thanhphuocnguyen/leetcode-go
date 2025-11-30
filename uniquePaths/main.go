package main

import "fmt"

func uniquePaths(m int, n int) int {
	// we know that first place have (m-1) and (n-1) mean (m+n-2) possible way to go to [m-1][n-1] pos
	// the combination formula between (m-1) and (n-1) moves is N!/(r!*(N-r))!
	// substitute N = (m+n-2) into formula we have (m+n-2)!/(r!*(m+n-2-r))!
	ans := 1
	if n > m {
		m, n = n, m
	}
	for i, j := m+n-2, 1; i >= m; i, j = i-1, j+1 {
		ans = (ans * i) / j
	}

	return ans
}

func main() {
	fmt.Println(uniquePaths(80, 2))
}
