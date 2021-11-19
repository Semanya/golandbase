package main

import (
	"time"
	"sync"
	"fmt"
	"context"
)

const (
	k1   = "key1"
	step = 7
)

type Cache struct {
	storage map[string]int
	mu         sync.RWMutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
	cache := Cache{storage: make(map[string]int)}
	semaphore := make(chan int, 5)
	L:
	for i := 0; i < 50; i++ {
		select {
		case semaphore <- i:
		    go func() {
			    defer func() {
			    	<-semaphore
			    }()
		    cache.Increase(k1, step)
		    fmt.Println(cache.Get(k1))
		    // time.Sleep(time.Millisecond * 10)
		}()
	    case <-ctx.Done():
			fmt.Println("go go break L")
			break L
		default:
			fmt.Println("Waiting")
			time.Sleep(time.Millisecond * 10)
	}
}
	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 10)
	}

	LL:
	for i := 0; i < 50; i++ {
	i := i // copy variable
	    select {
	    case semaphore <- i:
	    go func() {
		    defer func() {
		    	<-semaphore
		        }()
	 	cache.Set(k1, step*i)
	 	fmt.Println(cache.Get(k1))
	  	// time.Sleep(time.Millisecond * 100)
	}()
    case <-ctx.Done():
		fmt.Println("go go break LL")
	    break LL
    default:
	    time.Sleep(time.Millisecond * 10)
	}
}
	for len(semaphore) > 0 {
	  	time.Sleep(time.Millisecond * 10)
	  }

}