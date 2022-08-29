package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	m := make(map[int]int)
	var mut sync.Mutex
	var wg sync.WaitGroup
	var n = 10

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(iter int) {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			fmt.Println("work gorutine #", strconv.Itoa(iter))
			mut.Lock()
			m[iter] = iter * iter
			fmt.Println("map = ", m)
		}(i)
	}
	wg.Wait()

	fmt.Println("========================")
	fmt.Println("map = ", m)
}
