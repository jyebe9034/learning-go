package main

import (
	"fmt"

	"github.com/jyebe9034/golearning/accounts"
)

func main() {
	accounts := accounts.NewAccount("hannah")
	fmt.Println(accounts)
}
