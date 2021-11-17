package main

import (
	"time"
	"sync"
	"fmt"
)

type Cache struct {
	storage map[string]int
	mu         sync.Mutex
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
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

func main() {
	months := &Cache{
		storage: map[string]int{
			"Jan": 31,
			"Feb": 28,
			"Mar": 31,
		  },
	}
    go months.Increase("May", 7)
	months.Increase("Apr", 7)
	time.Sleep(time.Second)
	fmt.Println(months.Get("Apr"))
	fmt.Println(months.Get("May"))
}