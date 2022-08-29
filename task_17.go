package main

import (
	"fmt"
)

func binSearch(sl []float32, toFind float32) int { // возвращает индекс или -1, если не найдено
	left, right := 0, len(sl)-1
	mid := (len(sl) - 1) / 2

	for left <= right {
		mid = (right + left) / 2
		if sl[mid] == toFind {
			return mid

		} else if sl[mid] > toFind {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	sl := []float32{}

	for i := 0; i < 15; i++ {
		sl = append(sl, float32(i*i))
	}

	fmt.Println(sl)
	fmt.Println("================")
	fmt.Println("index find = ", binSearch(sl, 49))
	fmt.Println("index find = ", binSearch(sl, 63))

}
