package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	startTransaction(cust)

	fmt.Println(cust)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
