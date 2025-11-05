package main

import "fmt"

func update(tree []int, i, val int) {
	for i < len(tree) {
		tree[i] += val
		i = i + (i & -i)
	}
}

func querySum(tree []int, i int) int {
	sum := 0
	for i > 0 {
		sum += tree[i]
		i = i - (i & -i)
	}
	return sum
}

func processQueries(queries []int, m int) []int {
	n := len(queries)
	idxMp := make(map[int]int)
	fwTree := make([]int, n+m+1)
	for i := 1; i <= m; i++ {
		idxMp[i] = n + i
		update(fwTree, n+i, 1)
	}
	ans := make([]int, n)
	for i, q := range queries {
		ans[i] = querySum(fwTree, idxMp[q]-1)
		update(fwTree, idxMp[q], -1)
		update(fwTree, n, 1)
		idxMp[q] = n
		n--
	}
	return ans
}

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}
