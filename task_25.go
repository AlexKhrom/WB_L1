package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	now := time.Now().UnixNano()

	for {
		if now+int64(duration) <= time.Now().UnixNano() {
			break
		}
	}

}

func main() {
	fmt.Println("sleep for a few seconds  ", 5)
	fmt.Println("time now = ", time.Now().String())
	mySleep(time.Second * 5)
	fmt.Println("time after = ", time.Now().String())

}
