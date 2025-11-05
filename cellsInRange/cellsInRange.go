package main

import "fmt"

func cellsInRange(s string) []string {
	fromCol, toCol := s[0], s[3]
	fromRow, toRow := s[1], s[4]
	ans := make([]string, 0)
	for i := fromCol; i <= toCol; i++ {
		for j := fromRow; j <= toRow; j++ {
			ans = append(ans, string([]byte{i, j}))
		}
	}
	return ans
}

func main() {
	fmt.Println(cellsInRange("K1:L2"))
	fmt.Println(cellsInRange("A1:F1"))
}
