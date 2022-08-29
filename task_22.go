package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func sum(x, y *big.Float) *big.Float {
	return big.NewFloat(0).Add(x, y)
}

func subtr(x, y *big.Float) *big.Float {
	return big.NewFloat(0).Sub(x, y)
}

func multiply(x, y *big.Float) *big.Float {
	return big.NewFloat(0).Mul(x, y)
}

func div(x, y *big.Float) *big.Float {
	return big.NewFloat(0).Quo(x, y)
}

func main() {
	x := big.NewFloat(10000000000 * float64(rand.Int()))
	y := big.NewFloat(10000000000 * float64(rand.Int()))
	fmt.Println(x, y)

	fmt.Println(sum(x, y))
	fmt.Println(subtr(x, y))
	fmt.Println(multiply(x, y))
	fmt.Println(div(x, y))
}
