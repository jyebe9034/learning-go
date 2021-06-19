package accounts

type Accounts struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Accounts {
	account := Accounts{owner: owner, balance: 0}
	return &account
}
