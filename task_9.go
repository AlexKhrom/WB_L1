package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < len(sl); i++ {
			num := <-ch1
			ch2 <- num * 2
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(sl); i++ {
			fmt.Println(<-ch2)
		}
	}()

	for i := 0; i < len(sl); i++ {
		ch1 <- sl[i]
	}

	wg.Wait()

}
