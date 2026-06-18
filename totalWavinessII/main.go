package main

import "fmt"

func totalWaviness(num1 int64, num2 int64) int64 {
	return solve(num2) - solve(num1-1)
}

func solve(num int64) int64 {
	dpCnt := make([][][]int64, 16)
	dpSum := make([][][]int64, 16)
	digits := make([]int, 0)
	x := num
	for x > 0 {
		digits = append(digits, int(x%10))
		x /= 10
	}
	for i := range len(digits) / 2 {
		digits[i], digits[len(digits)-1-i] = digits[len(digits)-1-i], digits[i]
	}
	for i := range 16 {
		dpCnt[i] = make([][]int64, 11)
		dpSum[i] = make([][]int64, 11)
		for j := range 11 {
			dpCnt[i][j] = make([]int64, 11)
			dpSum[i][j] = make([]int64, 11)
			for k := range 11 {
				dpCnt[i][j][k] = -1
				dpSum[i][j][k] = -1
			}
		}
	}

	var dfs func(pos int, tight, started bool, lastDigit, secondLastDigit int) (total int64, totalWaviness int64)
	dfs = func(pos int, tight bool, started bool, lastDigit int, secondLastDigit int) (total int64, totalWaviness int64) {
		if pos == len(digits) {
			return 1, 0
		}
		if !tight && dpCnt[pos][lastDigit+1][secondLastDigit+1] != -1 {
			return dpCnt[pos][lastDigit+1][secondLastDigit+1], dpSum[pos][lastDigit+1][secondLastDigit+1]
		}

		limit := 9
		if tight {
			limit = int(digits[pos])
		}
		for d := 0; d <= limit; d++ {
			nextTight := tight && d == limit
			var childCnt, childSum int64 = 0, 0
			var extra int64 = 0
			if !started {
				if d == 0 {
					childCnt, childSum = dfs(pos+1, nextTight, false, -1, -1)
				} else {
					childCnt, childSum = dfs(pos+1, nextTight, true, d, -1)
				}

			} else if secondLastDigit == -1 {
				childCnt, childSum = dfs(pos+1, nextTight, true, d, lastDigit)
			} else {
				if lastDigit > secondLastDigit && lastDigit > d {
					extra = 1
				}
				if lastDigit < secondLastDigit && lastDigit < d {
					extra = 1
				}
				childCnt, childSum = dfs(pos+1, nextTight, true, d, lastDigit)
			}
			total += childCnt
			totalWaviness += (childSum + extra*childCnt)
		}

		if !tight && secondLastDigit >= 0 && lastDigit >= 0 {
			dpCnt[pos][lastDigit][secondLastDigit] = total
			dpSum[pos][lastDigit][secondLastDigit] = totalWaviness
		}

		return
	}
	_, total := dfs(0, true, false, -1, -1)
	return total
}

func main() {
	fmt.Println(totalWaviness(120, 130))
}
