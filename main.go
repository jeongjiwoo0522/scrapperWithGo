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
}