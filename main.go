package main

import (
	"fmt"

	"github.com/project/helloworld/accounts"
)

func main() {
	account := accounts.NewAccount("juno")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance(), account.Owner())
}
