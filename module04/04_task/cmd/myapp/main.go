package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

var (
	price int = 5000
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 400, true)

	fmt.Println(startTransactionDynamic(cust))

	cust.CalcDiscount()

	fmt.Printf("%+v\n", cust)

	rlyprice, err := internal.CalcPrice(cust, price)
	if err != nil {
		fmt.Printf("%+v\n %+v\n", rlyprice, err)
	} else {
		fmt.Printf("%+v\n", rlyprice)
	}
}

func startTransactionDynamic(w interface{}) error {
	_, ok := w.(internal.Discounter)
	if !ok {
		return errors.New("Incorrect type")
	}
	return errors.New("Type correct")
}
