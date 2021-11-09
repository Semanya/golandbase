package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Name     string
	Age      int
	balance  int
	debt     int
	discount bool
}

func (c *Customer) CalcDiscount() (int, error) {
	if !c.discount {
		return 0, errors.New("Discount not available")
	}
	result := DEFAULT_DISCOUNT - c.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		discount: discount,
	}
}

func CalcPrice(c Customer, cost int) (int, error) {
	var finalePrice int
	m, err := c.CalcDiscount()
	if err != nil {
		finalePrice = 0
	} else {
		finalePrice = cost - m
		if finalePrice < 0 {
			finalePrice = 0
		}
	}
	return finalePrice, err

}
