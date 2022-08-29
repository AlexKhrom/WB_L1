package main

import (
	"fmt"
	"sync"
)

func main() {
	mas := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var mut sync.Mutex
	var sum int
	wg.Add(len(mas))

	for _, dig := range mas {
		go func(num int) {
			defer wg.Done()
			fmt.Println("hi from ", num)
			mut.Lock()
			sum += num * num
			mut.Unlock()
		}(dig)
	}
	wg.Wait()

	fmt.Println("sum = ", sum)
}
