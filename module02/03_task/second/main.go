package main

import (
	"fmt"
)

func main() {
	Work_days := []string{"Mon", "Tue", "Wed", "Thur", "Fri"}
	Weekends := []string{"Sun", "Sat"}
	Week := append(Weekends, Work_days...)
	fmt.Println("Week ", Week)
}
