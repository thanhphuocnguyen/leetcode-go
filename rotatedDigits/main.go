package main

import "fmt"

func rotatedDigits(n int) int {
	// brute force
	ans := 0
	for i := 1; i <= n; i++ {
		x := i
		curr := 0
		for x > 0 {
			d := x % 10
			switch d {
			case 0:
				fallthrough
			case 1:
				fallthrough
			case 8:
				curr = curr*10 + d
			case 2:
				curr = curr*10 + 5
			case 5:
				curr = curr*10 + 2
			case 6:
				curr = curr*10 + 9
			case 9:
				curr = curr*10 + 6
			default:
				curr = -1
			}
			if curr == -1 {
				break
			}

			x /= 10
		}
		if curr != -1 && rev(i) != curr {
			ans++
		}
	}
	return ans
}

func rev(x int) int {
	ans := 0
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}

	return ans
}

func main() {
	fmt.Println(rotatedDigits(20))
}
