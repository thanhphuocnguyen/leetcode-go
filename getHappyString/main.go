package main

import "fmt"

func getHappyString(n int, k int) string {
	var result = ""
	var backtrack func(str string) int
	backtrack = func(currentStr string) int {
		if n == len(currentStr) {
			k--
			if k == 0 {
				result = currentStr
				return 0
			}
			return k
		}
		for _, b := range []byte("abc") {
			if len(currentStr) > 0 && currentStr[len(currentStr)-1] == b {
				continue
			}
			k = backtrack(currentStr + string(b))
			if k == 0 {
				return 0
			}
		}
		return k
	}
	backtrack("")
	return result
}

func main() {
	fmt.Println(getHappyString(1, 4))
}
