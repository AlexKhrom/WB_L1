package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan int)
	var N = 5
	defer close(ch)

	for i := 0; i < N; i++ {
		go func(iter int) {
			for {
				data, ok := <-ch
				if !ok {
					break
				}
				fmt.Println("worker #"+strconv.Itoa(iter)+" read - ", data)
			}
		}(i)
	}

	for i := 0; ; i++ {
		ch <- i
	}
}
