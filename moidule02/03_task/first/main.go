package main

import (
	"fmt"
)

func main() {
	Week := []string{"Sun", "Mon", "Tue", "Wed", "Thur", "Fri", "Sat"}
	Work_days := make([]string, 5)
	copy(Work_days, Week[1:6])
	fmt.Println("Work_days is ", Work_days)
	Week = append(Week[:1], Week[6:]...)
	fmt.Println("weekends ", Week)
}
