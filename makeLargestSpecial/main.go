package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}

func makeLargestSpecial(s string) string {
	cnt, i := 0, 0
	ans := make([]string, 0, len(s))
	for j := 0; j < len(s); j++ {
		if s[j] == '1' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ans = append(ans, "1"+makeLargestSpecial(s[i+1:j])+"0")
			i = j + 1
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	return strings.Join(ans, "")
}
