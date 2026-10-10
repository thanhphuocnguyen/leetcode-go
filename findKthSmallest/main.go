package main

import "fmt"

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)
	minCoin := coins[0]
	for i := 1; i < n; i++ {
		minCoin = min(minCoin, coins[i])
	}

	// bisec X
	var lo, hi int64 = 1, int64(k * minCoin)
	for lo < hi {
		mid := (lo + hi) / 2
		if count(mid, coins) >= int64(k) {
			hi = mid
		} else {
			lo = lo + 1
		}
	}

	return lo
}

func count(x int64, coins []int) int64 {
	var ans int64
	var dfs func(idx int, currLcm int64)

	dfs = func(idx int, currLcm int64) {
		for i := idx; i < len(coins); i++ {
			newLcm := lcm(currLcm, int64(coins[i]))
			if newLcm > x {
				continue
			}
			contrib := x / newLcm
			subsetSize := i - idx + 1
			if subsetSize%2 == 0 {
				ans -= contrib
			} else {
				ans += contrib
			}

			dfs(i+1, newLcm)
		}
	}

	dfs(0, 1)

	return ans
}

func lcm(a, b int64) int64 {
	if a == 0 && b == 0 {
		return 0
	}

	return (a / gcd(a, b)) * b
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	fmt.Println(findKthSmallest([]int{3, 6, 9}, 3))
}
