package main

import "fmt"

func getHappyString(n int, k int) string {
	var listStr = []string{}
	// var backtrack func(str string) int
	// backtrack = func(currentStr string) int {
	// 	if n == len(currentStr) {
	// 		k--
	// 		if k == 0 {
	// 			result = currentStr
	// 			return 0
	// 		}
	// 		return k
	// 	}
	// 	for _, b := range []byte("abc") {
	// 		if len(currentStr) > 0 && currentStr[len(currentStr)-1] == b {
	// 			continue
	// 		}
	// 		k = backtrack(currentStr + string(b))
	// 		if k == 0 {
	// 			return 0
	// 		}
	// 	}
	// 	return k
	// }
	backtrack(&listStr, []byte{}, n, k)
	if len(listStr) < k {
		return ""
	}
	return listStr[k-1]
}

func backtrack(listStr *[]string, bs []byte, n, k int) {
	if n == len(bs) {
		*listStr = append(*listStr, string(bs))
		return
	}
	for _, b := range []byte{'a', 'b', 'c'} {
		if len(bs) > 0 && bs[len(bs)-1] == b {
			continue
		}
		bs = append(bs, b)
		backtrack(listStr, bs, n, k)
		bs = bs[0 : len(bs)-1]
	}
}

func main() {
	fmt.Println(getHappyString(3, 9))
	fmt.Println(getHappyString(1, 3))
	fmt.Println(getHappyString(1, 4))
}
