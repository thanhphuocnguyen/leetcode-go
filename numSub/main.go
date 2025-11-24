package main

const MOD = 1e9 + 7

func numSub(s string) int {
	cnt := 0
	ans := 0
	for _, r := range s {
		if r == '1' {
			cnt += 1
		} else {
			ans += (cnt * (cnt + 1) / 2) % MOD
			cnt = 0
		}
	}
	ans += (cnt * (cnt + 1) / 2) % MOD
	return ans
}

func main() {
	numSub()
}
