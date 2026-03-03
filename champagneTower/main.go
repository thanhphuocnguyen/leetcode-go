package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	n := query_row + 1
	dp := make([][]float64, n)
	for i := range n {
		dp[i] = make([]float64, i+1)
	}
	dp[0][0] = float64(poured)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] > 1.0 {
				exceed := (dp[i][j] - 1.0) / 2.0
				dp[i][j] = 1.0
				if i+1 < n {
					dp[i+1][j] += exceed
					dp[i+1][j+1] += exceed
				}
			}
		}
	}
	return dp[query_row][query_glass]
}

func main() {
	fmt.Println(champagneTower(100000009, 33, 17))
}
