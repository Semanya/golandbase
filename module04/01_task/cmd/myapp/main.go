package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

var (
	price int = 100
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 300, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Printf("%+v\n", cust)

	rlyprice, err := internal.CalcPrice(*cust, price)
	if err != nil {
		fmt.Printf("%+v\n %+v\n", rlyprice, err)
	} else {
		fmt.Printf("%+v\n", rlyprice)
	}
}
