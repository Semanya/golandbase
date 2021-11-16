package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	go printNumbers(44)
	printNumbers(45)
}

func printNumbers(i int) (c string){
 		fibN := fib(i)
 		fmt.Printf("\rFibonacci(%d) = %d\n", i, fibN)
		return
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
