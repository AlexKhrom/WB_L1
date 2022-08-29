package main

import "fmt"

func changeBit(x int64, position int, bit int64) int64 { // position начинается с 0
	if bit != 0 && bit != 1 {
		fmt.Println("bit should be 0 or 1, got :", bit)
		return x
	}

	var myBit int64 = (x >> position) % 2
	x = x ^ (myBit << position)
	x = x | (bit << position)
	return x
}

func main() {
	var x int64 = 15

	fmt.Println("new x = ", changeBit(x, 1, 0))
}
