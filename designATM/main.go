package main

import "fmt"

type ATM struct {
	prices       [5]int
	bankNotesCnt []int
}

func Constructor() ATM {
	return ATM{
		prices:       [5]int{20, 50, 100, 200, 500},
		bankNotesCnt: make([]int, 5),
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, cnt := range banknotesCount {
		this.bankNotesCnt[i] += cnt
	}
}

func (this *ATM) Withdraw(amount int) []int {
	// 600
	x := amount
	ans := make([]int, 5)
	newState := make([]int, 5)
	copy(newState, this.bankNotesCnt)
	for i := 4; i >= 0; i-- {
		cnt := newState[i]
		// 500

		price := this.prices[i]
		if cnt > 0 && x/price > 0 {
			// 600/500 = 1
			quot := x / price
			// cnt = 1
			// taken = 1
			taken := min(quot, cnt)
			ans[i] = taken
			// cntRem = 0
			cntRem := quot - taken
			newState[i] -= taken
			// new x = 0*500 + (600%500) = 100
			x = cntRem*price + x%price
		}
		if x == 0 {
			break
		}

	}
	// mean still amount but there are no prices to pick
	if x > 0 {
		return []int{-1}
	}
	this.bankNotesCnt = newState
	return ans
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

func main() {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(atm.Withdraw(600))
	atm.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println(atm.Withdraw(600))
	fmt.Println(atm.Withdraw(550))
	fmt.Println()
}
