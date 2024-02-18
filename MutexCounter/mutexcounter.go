package main

import (
	"fmt"
	"sync"
	"time"
)

// safeCounter is safe to use concurrently
type safeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc icrements the counter for the given key
func (c *safeCounter) Inc(key string) {
	c.mu.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *safeCounter) value(key string) int {
	c.mu.Lock()
	//Lock so onely one goroutine at a time can access the map c.v. Defer ensures the mutex gets unlocked
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := safeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}
