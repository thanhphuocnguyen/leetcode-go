package main

import "fmt"

func numberOfWays(s string) int64 {
	// maintain 4 vars for ones, zeros on the left and right
	zerosRight, onesRight := 0, 0
	zerosLeft, onesLeft := 0, 0
	for _, r := range s {
		if r == '0' {
			zerosRight++
		} else {
			onesRight++
		}
	}
	var ans int64

	for _, r := range s {
		if r == '0' {
			// if encounter '0' we have combination for them is onesLeft*oneRight to have the good pairs
			ans += int64(onesLeft * onesRight)
			zerosLeft++
			zerosRight--
		} else {
			// if encounter '0' we have combination for them is onesLeft*oneRight to have the good pairs
			ans += int64(zerosLeft * zerosRight)
			onesLeft++
			onesRight--
		}
	}

	return ans
}

func main() {
	fmt.Println(numberOfWays("001101"))
}
