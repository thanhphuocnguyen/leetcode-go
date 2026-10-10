package main

import "fmt"

func minInsertions(s string) int {
	balance := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			balance++
		} else {
			if i+1 < len(s) && s[i] == s[i+1] {
				// case ()) ✅
				i += 1
			} else {
				// case () we need ): ())
				ans += 1
			}

			// --- case ()) or ))
			if balance == 0 {
				// case ) then add an (
				ans += 1
			} else {
				// case ()) then -1 in balance
				balance--
			}
		}
		// switch r {
		// case '(':
		// 	// case ()(.. | )(...
		// 	if rightCnt > 0 {
		// 		if balance == 0 {
		// 			ans += balance / 2
		// 			ans += 2 - (balance % 2)
		// 		} else {
		// 			ans++
		// 			balance--
		// 		}
		// 		rightCnt = 0
		// 	}
		// 	balance++
		// case ')':
		// 	rightCnt++
		// 	if rightCnt == 2 {
		// 		rightCnt -= 2
		// 		if balance == 0 {
		// 			ans++
		// 			rightCnt = 0
		// 		} else {
		// 			balance--
		// 		}
		// 	}
		// }
	}
	// if balance > 0 {
	// 	return ans + (balance * 2) - rightCnt
	// }
	// if rightCnt > 0 {
	// 	return ans + 2
	// }
	return ans + balance*2
}

func main() {
	// ())
	fmt.Println(minInsertions(")))))))"))
	fmt.Println(minInsertions("(()))(()))()())))"))
	fmt.Println(minInsertions("))())("))
	fmt.Println(minInsertions("())"))
	fmt.Println(minInsertions("(()))"))
}
