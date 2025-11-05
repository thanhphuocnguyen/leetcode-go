package main

const MOD = 1_000_000_007

func lengthAfterTransformations(s string, t int) int {
	freq := make([]int, 26)
	for _, c := range s {
		freq[int(c-'a')]++
	}
	for i := 0; i < t; i++ {
		next := make([]int, 26)
		next[0] = freq[25]
		next[1] = (freq[0] + freq[25]) % MOD
		for j := 2; j < 26; j++ {
			next[j] = freq[j-1]
		}
		freq = next
	}
	ans := 0
	for _, cnt := range freq {
		ans = (ans + cnt) % MOD
	}
	return ans
}

func main() {
	// Example usage
	s := "abcyy"
	t := 2
	result := lengthAfterTransformations(s, t)
	println(result) // Output: 3
}
