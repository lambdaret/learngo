package main

import (
	"fmt"

	"github.com/lambdaret/learngo/accounts"
)



func main() {
	account := accounts.NewAccount("nico")
	
	fmt.Println(account)
}