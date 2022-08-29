package main

import "fmt"

func main() {
	minus, plus := -1, 1
	fmt.Printf("minus = %d, plus = %d \n", minus, plus)

	minus, plus = plus, minus

	fmt.Printf("change - minus = %d, plus = %d \n", minus, plus)

	minus = plus + minus
	plus = minus - plus
	minus = minus - plus

	fmt.Printf("change - minus = %d, plus = %d \n", minus, plus)

}
