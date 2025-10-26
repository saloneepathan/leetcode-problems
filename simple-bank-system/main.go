package main

import (
	"fmt"
)

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (this *Bank) isValidAccount(account int) bool {
	return account >= 1 && account <= len(this.balance)
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.isValidAccount(account1) || !this.isValidAccount(account2) {
		return false
	}
	if this.balance[account1-1] < money {
		return false
	}
	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.isValidAccount(account) {
		return false
	}
	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isValidAccount(account) {
		return false
	}
	if this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}

func main() {
	bank := Constructor([]int64{10, 100, 20, 50, 30})

	fmt.Println(bank.Withdraw(3, 10))    // true
	fmt.Println(bank.Transfer(5, 1, 20)) // true
	fmt.Println(bank.Deposit(5, 20))     // true
	fmt.Println(bank.Transfer(3, 4, 15)) // false
	fmt.Println(bank.Withdraw(10, 50))   // false

	fmt.Println("Final balances:", bank.balance)
}

/**
Expected Output:
true
true
true
false
false
Final balances: [30 100 10 50 30]
*/
