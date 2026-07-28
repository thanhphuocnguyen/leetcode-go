package main

import (
	"fmt"
)

func sumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1_000_000_007
	n := len(s)
	prefixSum := make([]int, n+1)
	cnt := make([]int, n+1)
	pow10 := make([]int, n+1)
	pow10[0] = 1

	P := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		c := s[i-1]
		if c == '0' {
			prefixSum[i] = prefixSum[i-1]
			P[i] = P[i-1]
			cnt[i] = cnt[i-1]
		} else {
			num := int(c - '0')
			prefixSum[i] = prefixSum[i-1] + num
			P[i] = (P[i-1]*10 + int64(num)) % MOD
			cnt[i] = cnt[i-1] + 1
		}
		pow10[i] = pow10[i-1] * 10 % MOD
	}
	/*
		If we observe carefully, in terms of numbers, substring of 3, 7 actually means (substring of (0,7) - substring of (0,3)*(10^(7-3)).
		  Here, it will be:
		  98765432-98700000

		  This means, we ultimately just need to find (98765432-98700000)%(10^9+7), and then multiply it again with the sum, and then again mod it with the same (10^9+7) value.
		  So if 10^9+7 = MOD, the above mod operation becomes:
		  ((98765432 % MOD) - ((987 % MOD)*(10^(7-3) % MOD)) % MOD) % MOD
	*/
	ans := make([]int, len(queries))

	for i, q := range queries {
		l, r := q[0], q[1]+1
		x := (P[r] % MOD)
		offset := cnt[r] - cnt[l]
		sum := int64(prefixSum[r] - prefixSum[l])
		x -= P[l] * int64(pow10[offset])
		x %= MOD
		if x < 0 {
			x += MOD
		}
		ans[i] = int(x * sum % MOD)
	}

	return ans
}

func main() {
	fmt.Println(sumAndMultiply("42", [][]int{{0, 0}, {0, 1}, {1, 1}}))
	fmt.Println(sumAndMultiply("10203004", [][]int{{0, 7}, {1, 3}, {4, 6}}))
}
