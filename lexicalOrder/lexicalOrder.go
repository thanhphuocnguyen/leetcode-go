package main

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	curr := 1
	for i := 0; i < n; i++ {
		ans[i] = curr
		// 1 -> 10, 11; 2 -> 20, 21...
		if curr*10 < n {
			curr *= 10
		} else {
			// 19 -> 2, 29 -> 3
			if curr%10 == 9 || curr >= n {
				curr /= 10
			}
			// 10 11 12 13 14 15 16 17 18 19
			curr++
		}
	}
	return ans
}
