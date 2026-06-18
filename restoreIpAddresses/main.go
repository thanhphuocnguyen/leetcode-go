package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	// backtrack will go from traversal through string and keep track current dot number
	// beside it we track current idx and prev idx to make sure current number not have any leading zeros
	// if prev idx is '0' we should cut it, max len if current seg must not be larger than 3
	// if prev
	n := len(s)
	ans := make([]string, 0)
	path := []string{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) > 4 {
			return
		}
		if start >= n {
			if len(path) == 4 {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}

		for end := start + 1; end-start < 4 && end <= n; end++ {
			str := s[start:end]
			if len(str) > 1 && str[0] == '0' {
				return
			}
			num, _ := strconv.Atoi(str)
			if num <= 255 {
				path = append(path, str)
				backtrack(end)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}

	backtrack(0)
	return ans
}
func main() {
	fmt.Println(restoreIpAddresses("172162541"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("25525511135"))
}
