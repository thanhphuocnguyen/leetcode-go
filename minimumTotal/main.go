package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][0] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][len(triangle[i-1])-1] + triangle[i][j]
			} else {
				triangle[i][j] = triangle[i][j] + min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}

	ans := triangle[n-1][0]
	for _, val := range triangle[n-1] {
		ans = min(ans, val)
	}
	return ans
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
