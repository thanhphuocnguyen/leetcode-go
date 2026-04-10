package main

import "fmt"

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	strLen := (2 << n) - 1

	mid := (strLen / 2) + 1
	if mid == k {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		found := findKthBit(n-1, strLen-k+1)
		if found == '0' {
			return '1'
		} else {
			return '0'
		}
	}
	// s := []byte{'0'}
	// for i := 1; i <= n; i++ {
	// 	s = append(s, '1')
	// 	n := len(s)
	// 	// minus one to exclude '1'
	// 	for j := n - 2; j >= 0; j-- {
	// 		var b byte
	// 		if s[j] == '1' {
	// 			b = '0'
	// 		} else {
	// 			b = '1'
	// 		}
	// 		s = append(s, b)
	// 	}
	// }
	// return s[k]
}

func main() {
	fmt.Println(findKthBit(3, 1))
}
