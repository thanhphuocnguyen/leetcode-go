package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	freq := make([]int, 26)
	for _, r := range word {
		freq[r-'a']++
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	cnt := 1
	ans := 0
	for i := range 26 {
		if freq[i] == 0 {
			break
		}
		ans += cnt * freq[i]
		if i >= cnt*8-1 {
			cnt++
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
}
