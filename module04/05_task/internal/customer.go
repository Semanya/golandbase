package internal

type Customer struct {
	Name string
	Age  int
	Debtor
}

type Overduer struct {
	balance int
	debt    int
}

func (c *Overduer) WrOffDebt() error {
	c.debt = 0
	return nil
}

func NewOverduer(balance int, debt int) *Overduer {
	return &Overduer{
		balance: balance,
		debt:    debt}

}
