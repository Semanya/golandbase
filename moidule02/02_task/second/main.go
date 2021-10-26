package main

import (
	"fmt"
	"math"
)

func main() {
	var Radius = 35 / (2 * 3.14)
	var R *float64 = &Radius
	var S = 3.14 * (math.Pow(*R, 2))
	fmt.Printf("%.2f", S)
}
