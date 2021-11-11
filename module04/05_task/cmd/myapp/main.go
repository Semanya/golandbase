package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	over := internal.NewOverduer(10000, 1000)
	cust := internal.Customer{
		Name:   "Dmitry",
		Age:    23,
		Debtor: over}

	startTransaction(cust.Debtor)

	fmt.Println(cust.Debtor)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
