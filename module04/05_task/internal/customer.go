package internal

type Customer struct {
	Name     string
	Age      int
	discount bool
	*Overduer
}

type Overduer struct {
	balance int
	debt    int
}

func (c Overduer) WrOffDebt() error {
	c.debt = 0
	return nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name: name,
		Age:  age,
		Overduer: &Overduer{
			balance: balance,
			debt:    debt},
		discount: discount,
	}

}
