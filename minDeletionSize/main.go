package main

import "fmt"

func minDeletionSize(strs []string) int {
	ans := 0
	n := len(strs)
	sorted := make([]bool, n)
	m := len(strs[0])
	for j := 0; j < m; j++ {
		deleted := false
		for i := 1; i < n; i++ {
			if !sorted[i] && strs[i-1][j] > strs[i][j] {
				deleted = true
				ans++
				break
			}
		}
		if deleted {
			continue
		}
		for i := 1; i < n; i++ {
			if !sorted[i] && strs[i-1][j] < strs[i][j] {
				sorted[i] = true
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minDeletionSize([]string{"xga", "xfb", "yfa"}))
	fmt.Println(minDeletionSize([]string{"abx", "agz", "bgc", "bfc"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
	fmt.Println(minDeletionSize([]string{"xc", "yb", "za"}))
	fmt.Println(minDeletionSize([]string{"ca", "bb", "ac"}))
}
