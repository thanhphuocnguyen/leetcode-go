package main

import "fmt"

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range n {
		ans[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++
		// if top == bottom {
		// 	break
		// }

		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--

		// if right == left {
		// 	break
		// }

		for i := right; i >= left; i-- {
			ans[bottom][i] = num
			num++
		}

		bottom--
		// if top == bottom {
		// 	break
		// }
		for i := bottom; i >= top; i-- {
			ans[i][left] = num
			num++
		}
		left++
	}
	return ans
}

func main() {
	fmt.Println(generateMatrix(4))
}
