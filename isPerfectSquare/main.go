package main

import "fmt"

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		rm := num / mid
		if rm == mid {
			return true
		} else if mid > rm {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
}
