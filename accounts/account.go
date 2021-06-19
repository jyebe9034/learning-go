package accounts

import "errors"

// Accounts struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw cuz you're poor")

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // (a Account)는 리시버라고 불림
	a.balance += amount
}

// Balance of your account
func (a Account) Balnace() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
