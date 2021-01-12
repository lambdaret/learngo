package main

import (
	"fmt"
	"log"

	"github.com/lambdaret/learngo/accounts"
)



func main() {

	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance())
}