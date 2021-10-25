package main

import (
	"fmt"
	//	"log"
	"strconv"
)

var (
	first  string = "104"
	second int    = 35
)

func main() {
	first, _ := strconv.Atoi(first)
	second := strconv.Itoa(second)
	fmt.Print("first data is ", first)
	fmt.Printf(" and has type:  %T\n", first)
	fmt.Print("second data is ", second)
	fmt.Printf(" and has type: %T\n", second)

}
