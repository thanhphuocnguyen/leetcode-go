package main

import "fmt"

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	points1, points2 := make([][]int, 0, n), make([][]int, 0, n)
	for i := range n {
		for j := range n {
			if img1[i][j] == 1 {
				points1 = append(points1, []int{i, j})
			}

			if img2[i][j] == 1 {
				points2 = append(points2, []int{i, j})
			}

		}
	}
	cnt := make([][]int, 2*n)
	for i := range 2 * n {
		cnt[i] = make([]int, 2*n)
	}
	ans := 0
	for _, p1 := range points1 {
		for _, p2 := range points2 {
			dx, dy := p2[0]-p1[0]+n, p2[1]-p1[1]+n
			cnt[dx][dy]++
			ans = max(ans, cnt[dx][dy])
		}
	}
	return ans
}

func main() {
	fmt.Println(largestOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
}
