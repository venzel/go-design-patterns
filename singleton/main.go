package main

import (
	"dpsingleton/singleton"
	"fmt"
)

func main() {
	instance := singleton.GetInstance()

	customers := instance.CustomersRepository.ListAll()
	fmt.Println(customers)

	accounts := instance.AccountsRepository.ListAll()
	fmt.Println(accounts)
}
