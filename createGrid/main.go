package main

import "fmt"

func createGrid(m int, n int) []string {
	ans := make([]string, m)
	bs := make([]byte, n)
	for i := range m {
		for j := range n {
			if j == 0 || i == m-1 {
				bs[j] = '.'
			} else {
				bs[j] = '#'
			}
		}
		ans[i] = string(bs)
	}

	return ans
}

func main() {
	fmt.Println(createGrid(2, 3))
}
