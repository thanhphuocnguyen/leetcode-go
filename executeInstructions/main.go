package main

func executeInstructions(n int, startPos []int, s string) []int {
	m := len(s)
	ans := make([]int, m)
	for i := range m {
		startRow, startCol := startPos[0], startPos[1]
		cnt := 0
		for j := i; j < m; j++ {
			switch s[j] {
			case 'L':
				startCol--
			case 'R':
				startCol++
			case 'U':
				startRow--
			case 'D':
				startRow++
			}
			if !isValid(startRow, startCol, n) {
				break
			} else {
				cnt++
			}
		}
		ans[i] = cnt
	}
	return ans
}

func isValid(i, j, n int) bool {
	return i >= 0 && j >= 0 && i < n && j < n
}
