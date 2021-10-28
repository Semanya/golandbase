package main

import (
	"fmt"
)

type (
	AmericanVelocity float64
	EuropeanVelocity float64
)

var (
	first_metr  = 120.4
	second_metr = 130.0
)

func main() {
	first_kilometr := EuropeanVelocity((first_metr * 18) / 5)
	fmt.Printf("%.2f", first_kilometr)
	fmt.Printf(" and has type:  %T\n", first_kilometr)
	second_milya := AmericanVelocity((second_metr * 18) / (5 * 1.609))
	fmt.Printf("%.2f", second_milya)
	fmt.Printf(" and has type:  %T\n", second_milya)
}
