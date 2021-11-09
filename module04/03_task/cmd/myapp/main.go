package main

import (
	"fmt"
	"myapp/internal"
)

var (
	price int = 5000
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 200, true)

	cust.CalcDiscount()

	fmt.Printf("%+v\n", cust)

	rlyprice, err := internal.CalcPrice(cust, price)
	if err != nil {
		fmt.Printf("%+v\n %+v\n", rlyprice, err)
	} else {
		fmt.Printf("%+v\n", rlyprice)
	}
}
