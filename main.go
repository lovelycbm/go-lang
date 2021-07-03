package main

import (
	"fmt"

	"example.com/hello/accounts"
)

func main() {
	account := accounts.NewAccount("bong")

	account.Deposit(100)

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
