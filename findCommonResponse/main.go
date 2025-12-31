package main

import "fmt"

func findCommonResponse(responses [][]string) string {
	mp := make(map[string]int)
	for _, rsp := range responses {
		set := make(map[string]bool)
		for _, msg := range rsp {
			if _, ok := set[msg]; !ok {
				mp[msg]++
				set[msg] = true
			}
		}
	}
	currMax := 0
	ans := ""
	for k, v := range mp {
		if v > currMax {
			currMax = v
			ans = k
		} else if v == currMax {
			isEqual := true
			for i := range min(len(ans), len(k)) {
				if k[i] > ans[i] {
					isEqual = false
					break
				}
				if k[i] < ans[i] {
					isEqual = false
					ans = k
				}
			}
			if isEqual && len(k) < len(ans) {
				ans = k
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findCommonResponse([][]string{{"gzdk", "l", "l", "opo", "ny"}}))
	fmt.Println(findCommonResponse([][]string{{"good", "ok", "good"}, {"ok", "bad"}, {"bad", "notsure"}, {"great", "good"}}))
	fmt.Println(findCommonResponse([][]string{{"good", "ok", "good", "ok"}, {"ok", "bad", "good", "ok", "ok"}, {"good"}, {"bad"}}))
}
