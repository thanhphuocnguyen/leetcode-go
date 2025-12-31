package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	prefixSum := make([][]int, m)
	for i := range m {
		prefixSum[i] = make([]int, n+1)
		for c := 1; c <= n; c++ {
			prefixSum[i][c] = prefixSum[i][c-1] + matrix[i][c-1]
		}
	}
	ans := 0
	for c1 := 0; c1 <= n; c1++ {
		for c2 := c1 + 1; c2 <= n; c2++ {
			mp := make(map[int]int)
			mp[0] = 1
			sum := 0
			for r := 0; r < m; r++ {
				sum += prefixSum[r][c2] - prefixSum[r][c1]
				if val, ok := mp[sum-target]; ok {
					ans += val
				}
				mp[sum]++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0))
}
