package main

import (
	"fmt"
)

var (
	A *int
	B int = 5
)

func main() {
	A = &B
	fmt.Println(*A)
	*A = 10
	fmt.Println(B)
}
