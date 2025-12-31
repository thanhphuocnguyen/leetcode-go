package main

func maximumProfit(prices []int, k int) int64 {

	memo := make([][][]int, len(prices))
	for i := range len(prices) {
		memo[i] = make([][]int, k+1)
		for j := range k + 1 {
			memo[i][j] = make([]int, 3)
		}
	}
}

func dp(prices []int, k int, idx int, transactionsDone bool, transactionType, string, isTransactionRunning bool) int {

}
func main() {

}
