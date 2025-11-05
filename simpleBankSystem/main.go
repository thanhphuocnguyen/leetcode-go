package main

type Bank struct {
	accounts []int64
	n        int
}

func Constructor(balance []int64) Bank {
	accounts := make([]int64, len(balance)+1)
	for i, b := range balance {
		accounts[i+1] = b
	}
	return Bank{
		accounts: accounts,
		n:        len(balance) + 1,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > this.n || account2 > this.n {
		return false
	}
	if this.accounts[account1] < money {
		return false
	}
	this.accounts[account1] -= money
	this.accounts[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > this.n {
		return false
	}
	this.accounts[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > this.n {
		return false
	}
	if this.accounts[account] < money {
		return false
	}
	this.accounts[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
