package main

import (
	"fmt"
	"sync"
)

func main() {
	mas := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(mas))

	for _, dig := range mas {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(dig)
	}
	wg.Wait()

}
