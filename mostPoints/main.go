package main

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	memo := make([]int64, n)
	for i := range n {
		memo[i] = -1
	}
	return dp(memo, questions, 0)
}

func dp(memo []int64, questions [][]int, currentIdx int) int64 {
	if currentIdx >= len(questions) {
		return 0
	}
	if memo[currentIdx] != -1 {
		return memo[currentIdx]
	}
	includedQuestion := int64(questions[currentIdx][0]) + dp(memo, questions, currentIdx+questions[currentIdx][1]+1)
	excludedQuestion := dp(memo, questions, currentIdx+1)
	result := max(includedQuestion, excludedQuestion)
	memo[currentIdx] = result
	return result
}

func main() {
	println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
}
