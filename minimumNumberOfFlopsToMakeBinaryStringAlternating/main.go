package main

func minFlips(s string) int {
	n := len(s)
	// create s+s to easy to use sliding window to simulate flip operation
	str := s + s
	flipCnt := 0
	ans := n
	// using sliding window to run each window with len = n
	for r := range str {
		var currByte byte
		if r%2 == 0 {
			currByte = '1'
		} else {
			currByte = '0'
		}
		// flip '1010101'
		if str[r] != currByte {
			flipCnt++
		}

		if r >= n {
			// remove flip from window
			l := r - n
			var leftByte byte
			if l%2 == 0 {
				leftByte = '1'
			} else {
				leftByte = '0'
			}
			if str[l] != leftByte {
				flipCnt--
			}
		}

		if r >= n-1 {
			ans = min(ans, flipCnt, n-flipCnt)
		}
	}
	return ans
}
