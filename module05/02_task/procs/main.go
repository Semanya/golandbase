package main

import (
	"fmt"
	"runtime"
)

func main () {
	runtime.GOMAXPROCS(3)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Starting goroutine %d \n", i +1)
			for {}
		}(i)
	}

	select {}
}
