package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	var backtrack func(path []int, used int, k, n int)
	backtrack = func(path []int, used int, k, n int) {
		if k == 0 {
			if n == 0 {
				cp := make([]int, len(path))
				copy(cp, path)
				ans = append(ans, cp)
			} else {
				return
			}
		}

		for i := used + 1; i <= 9; i++ {
			if n-i >= 0 {
				path = append(path, i)
				backtrack(path, i, k-1, n-i)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	backtrack([]int{}, 0, k, n)
	return ans
}

func main() {
	fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(3, 7))
}
