package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	// mp := make(map[string]struct{})
	// i, j := 0, 0
	// for j < len(s) {
	// 	if j-i == k {
	// 		substr := s[i:j]
	// 		mp[substr] = struct{}{}
	// 		i++
	// 	}
	// 	j++
	// }
	// mp[s[i:j]] = struct{}{}
	// numOfSubstr := 1 << k
	// return len(mp) == numOfSubstr

	// using rolling hash
	// all of the codes in k is 1 << k (2^k)
	n := 1 << k
	remains := n
	bins := make([]bool, n)
	// take mask of k ex: 10000 -> 01111
	mask := n - 1
	currHash := 0
	if len(s)-k+1 < n {
		return false
	}
	// get first hash window
	for i := range k {
		currHash = (currHash << 1) | int(s[i]-'0')
	}
	remains--
	bins[currHash] = true
	// rolling hash begins
	// calculate next hash by using sliding window technique
	for i := k; i < len(s); i++ {
		currHash = ((currHash << 1) | int(s[i]-'0')) & mask
		if !bins[currHash] {
			remains--
		}
		bins[currHash] = true
	}
	return remains == 0
}

func main() {
	fmt.Println(hasAllCodes("0", 20))
	fmt.Println(hasAllCodes("00000000001011100", 3))
	fmt.Println(hasAllCodes("00110", 2))
}
