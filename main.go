package main

import (
	"fmt"
	"github/jeongjiwoo0522/scrapperWithGo/account"
)

func main() {
	account := account.NewAccount("nico")
	fmt.Println(*account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(12)
	if err != nil {
		fmt.Println(err)
	}
	account.ChangeOwner("jiwoo")
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}