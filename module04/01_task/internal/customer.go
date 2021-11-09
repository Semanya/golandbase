package internal

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
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
