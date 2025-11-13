package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	memo := make([][][]int, len(strs))
	for i := range len(strs) {
		memo[i] = make([][]int, m+1)
		for j := range m + 1 {
			memo[i][j] = make([]int, n+1)
		}
	}
	ans := dp(strs, memo, 0, m, n, 0, 0)
	return ans
}

func dp(strs []string, memo [][][]int, idx, m, n, curZeros, curOnes int) int {
	if curZeros == m && curOnes == n || idx == len(strs) {
		return 0
	}

	if memo[idx][curZeros][curOnes] != 0 {
		return memo[idx][curZeros][curOnes]
	}

	// pick or skip
	zeros, ones := 0, 0
	for _, r := range strs[idx] {
		if r == '0' {
			zeros++
		} else {
			ones++
		}
	}
	maxLen := 0
	if curZeros+zeros <= m && curOnes+ones <= n {
		// pick
		maxLen = 1 + dp(strs, memo, idx+1, m, n, curZeros+zeros, curOnes+ones)
	}

	// skip
	notPick := dp(strs, memo, idx+1, m, n, curZeros, curOnes)
	memo[idx][curZeros][curOnes] = max(notPick, maxLen)
	return memo[idx][curZeros][curOnes]
}

func main() {
	fmt.Println(findMaxForm([]string{"0", "1101", "01", "00111", "1", "10010", "0", "0", "00", "1", "11", "0011"}, 63, 36))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
