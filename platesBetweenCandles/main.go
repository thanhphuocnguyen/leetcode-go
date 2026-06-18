package main

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
	prefixSum := make([]int, len(s)+1)

	candleIndices := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			prefixSum[i+1] = prefixSum[i]
			candleIndices = append(candleIndices, i)
		} else {
			prefixSum[i+1] = prefixSum[i] + 1
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		from, to := q[0], q[1]
		l := lowerBound(candleIndices, from)
		r := upperBound(candleIndices, to)
		if l == -1 || r == -1 {
			ans[i] = 0
			continue
		}
		if l < r {
			ans[i] = prefixSum[candleIndices[r]] - prefixSum[candleIndices[l]]
		} else {
			ans[i] = 0
		}

	}
	return ans
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

func upperBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] > target {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}

	return ans
}

func main() {
	fmt.Println(platesBetweenCandles("|*|*", [][]int{{3, 3}}))
	fmt.Println(platesBetweenCandles("*|*|||", [][]int{{0, 0}, {1, 3}}))
	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
}
