package main

import "fmt"

func main() {
	var sl = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var i = 6
	fmt.Println("slice = ", sl)
	sl = append(sl[:i], sl[i+1:]...)
	fmt.Printf("slice after delete sl[%d] elem = %v", i, sl)
	fmt.Println()

}
