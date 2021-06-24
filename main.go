package main

import (
	"fmt"

	"github.com/jyebe9034/golearning/accounts"
)

func main() {
	account := accounts.NewAccount("hannah")
	account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(account.Balnace(), account.Owner())
	fmt.Println(account)
}
