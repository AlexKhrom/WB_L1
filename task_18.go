package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mut   sync.Mutex
	count int
}

func (c *Counter) inc() {
	defer c.mut.Unlock()
	c.mut.Lock()
	c.count++
}

func (c *Counter) dec() {
	defer c.mut.Unlock()
	c.mut.Lock()
	c.count--
}

func (c *Counter) getCount() int {
	return c.count
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	var n = 1000
	var badCounter int

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			badCounter = badCounter + 1
			counter.inc()
		}()
	}
	wg.Wait()

	fmt.Println("my counter = ", counter.getCount())
	fmt.Println("bad counter = ", badCounter)

}
