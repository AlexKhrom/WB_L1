package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds = 5
	ch := make(chan int)

	startTime := time.Now().Unix()

	go func() {
		for val := range ch {
			fmt.Println("read - ", val)
		}
	}()

	for i := 0; ; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 200)
		if time.Now().Unix() >= startTime+int64(seconds) {
			close(ch)
			break
		}
	}
}
