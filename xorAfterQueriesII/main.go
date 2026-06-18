package main

import "math"

const MOD = 1_000_000_007

func xorAfterQueries(nums []int, queries [][]int) int {
	n := len(nums)
	T := int(math.Sqrt(float64(n)))
	groups := make([][][]int, T)
	for i := 0; i < T; i++ {
		groups[i] = make([][]int, 0)
	}

	for _, q := range queries {
		l, r := q[0], q[1]
		k, v := q[2], q[3]
		if k < T {
			groups[k] = append(groups[k], []int{l, r, v})
		} else {
			for i := l; i <= r; i += k {
				nums[i] = int(int64(nums[i]*v) % MOD)
			}
		}
	}

	// Difference array
	diff := make([]int64, n+T)
	for k := 1; k < T; k++ {
		if len(groups[k]) == 0 {
			continue
		}
		for i := range n + T {
			diff[i] = 1
		}

		for _, q := range groups[k] {
			l, r, v := q[0], q[1], q[2]
			diff[l] = (diff[l] * int64(v)) % MOD
			R := l + ((r-l)/k+1)*k
			diff[R] = (diff[R] * int64(pow(int64(v), MOD-2))) % MOD
		}

		for i := k; i < n; i++ {
			diff[i] = (diff[i] * diff[i-k]) % MOD
		}
		for i := range n {
			nums[i] = int((int64(nums[i]) * diff[i]) % MOD)
		}
	}
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

func pow(a, b int64) int {
	var ans int64 = 1
	for b > 0 {
		if b&1 == 1 {
			ans = (ans * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return int(ans)
}
