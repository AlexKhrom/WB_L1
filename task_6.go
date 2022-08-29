package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	mut.Lock()
	wg.Add(1)

	go func() {
		fmt.Println("start executing gorutine 1") // останавливает выполнение чтение из канала
		<-ch
		fmt.Println("end executing gorutine 1")

	}()
	go func() {
		fmt.Println("start executing gorutine 2") // останавливает выполнение запись в канал
		ch2 <- 1
		fmt.Println("end executing gorutine 2")

	}()
	go func() {
		fmt.Println("start executing gorutine 3") // останавливает выполнение семафор
		mut.Lock()
		mut.Unlock()
		fmt.Println("end executing gorutine 3")
	}()

	go func() {
		fmt.Println("start executing gorutine 4") // останавливает выполнение time.Sleep
		time.Sleep(time.Second * 2)
		fmt.Println("end executing gorutine 4")
	}()

	go func() {
		fmt.Println("start executing gorutine 5") // останавливает выполнение select
		for {
			select {
			case <-ch3:
				fmt.Println("end executing gorutine 5")
				return
			}
		}
	}()

	go func() {
		fmt.Println("start executing gorutine 6") // останавливает выполнение time.Sleep
		wg.Wait()
		fmt.Println("end executing gorutine 6")
	}()

	time.Sleep(time.Second * 4)
	ch <- 1
	<-ch2
	ch3 <- 1
	mut.Unlock()
	wg.Done()

	time.Sleep(time.Second)

}
