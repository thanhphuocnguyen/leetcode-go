package main

import "fmt"

func numberOfWays(corridor string) int {
	const MOD int = 1_000_000_007
	// save position of seats as splitted segments
	posArr := make([]int, 0)
	for i, obj := range corridor {
		if obj == 'S' {
			posArr = append(posArr, i)
		}
	}
	// if number of seats is not odd so there are no ways to find
	if len(posArr) == 0 || len(posArr)%2 != 0 {
		return 0
	}

	ans := 1
	for i := 2; i < len(posArr); i += 2 {
		segment := posArr[i] - posArr[i-1]
		// find product of current ways and the current segment to find answer
		ans = (ans * segment) % MOD
	}

	return ans
}

func main() {
	fmt.Println(numberOfWays("SSPPSPS"))
}
