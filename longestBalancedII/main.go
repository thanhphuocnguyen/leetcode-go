package main

import "fmt"

func longestBalanced(s string) int {
	x := findOne(s)
	y := max(findPair(s, 'a', 'b'), findPair(s, 'b', 'c'), findPair(s, 'c', 'a'))
	z := findTriple(s)
	return max(x, y, z)
}

func findOne(s string) int {
	n := len(s)
	ans := 1
	cnt := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			ans = max(cnt, ans)
		}
	}
	return ans
}

func findPair(s string, x, y rune) int {
	cnt := 0
	ans := 0
	mp := make(map[int]int)
	mp[0] = -1
	for i, r := range s {
		if r == x {
			cnt++
		} else {
			cnt--
		}
		if idx, ok := mp[cnt]; ok {
			ans = max(ans, i-idx)
		} else {
			mp[cnt] = i
		}
	}
	return ans
}

func findTriple(s string) int {
	cntA, cntB, cntC := 0, 0, 0
	mp := make(map[[2]int]int)
	ans := 0
	for i, r := range s {
		switch r {
		case 'a':
			cntA++
		case 'b':
			cntB++
		default:
			cntC++
		}
		key := [2]int{cntB - cntA, cntC - cntA}
		if val, ok := mp[key]; ok {
			ans = max(ans, i-val)
		} else {
			mp[key] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalanced("ccaca"))
	fmt.Println(longestBalanced("abbac"))
	fmt.Println(longestBalanced("aabcc"))
	fmt.Println(longestBalanced("aba"))
}
