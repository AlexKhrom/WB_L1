package main

import (
	"fmt"
	"math/rand"
)

func quickSort(sl []float32) {
	if len(sl) < 2 {
		return
	} else {
		left, right := 0, len(sl)-1
		med := len(sl) / 2

		sl[med], sl[right] = sl[right], sl[med]

		for i := range sl {
			if sl[i] < sl[right] {
				sl[left], sl[i] = sl[i], sl[left]
				left++
			}
		}

		sl[left], sl[right] = sl[right], sl[left]

		quickSort(sl[:left])
		quickSort(sl[left+1:])
		return
	}

}

func main() {
	sl := []float32{}

	for i := 0; i < 10; i++ {
		sl = append(sl, float32(rand.Intn(100)))
	}
	fmt.Println(sl)
	quickSort(sl)
	fmt.Println(sl)

}
