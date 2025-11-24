package main

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
	diff := make([][]int, n)
	ans := make([][]int, n)

	for i := range n {
		diff[i] = make([]int, n)
		ans[i] = make([]int, n)
	}
	// use difference array technique to update each row
	for _, query := range queries {
		row1, row2 := query[0], query[2]
		col1, col2 := query[1], query[3]
		for i := row1; i <= row2; i++ {
			update(diff[i], col1, col2)
		}
	}
	for i := range n {
		ans[i] = append(ans[i], diff[i][0])
		for j := 1; j < n; j++ {
			diff[i][j] += diff[i][j-1]
			ans[i] = append(ans[i], diff[i][j])
		}
	}
	return ans
}

func update(arr []int, l, r int) {
	arr[l]++
	if r+1 < len(arr) {
		arr[r+1]--
	}
}

func main() {
	fmt.Println(rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
}
