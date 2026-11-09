package main

import "fmt"

// min operations
func minOperations(s string) int {
	n := len(s)
	isSorted := true

	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			isSorted = false
		}
	}
	if isSorted {
		return 0
	}

	if n == 2 {
		return -1
	}
	mnCnt, mxCnt := 0, 0
	mn, mx := 0, n-1
	for i := 1; i < n; i++ {
		if s[i] < s[mn] {
			mn = i
			mnCnt = 1
		} else if s[i] == s[mn] {
			mnCnt++
		}
	}

	for i := n - 2; i >= 0; i-- {
		if mx == -1 {
			mx = i
		}
		if s[i] > s[mx] {
			mx = i
			mxCnt = 1
		} else if s[i] == s[mx] {
			mxCnt++
		}
	}

	if mn == 0 || mx == n-1 {
		return 1
	}

	if mn == n-1 && mx == 0 && mxCnt == 1 && mnCnt == 1 {
		return 3
	}

	return 2
}

func main() {
	fmt.Println(minOperations("dog"))
	fmt.Println(minOperations("srrp"))
	fmt.Println(minOperations("edc"))
}
