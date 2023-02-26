package main

import (
	"fmt"
	"math/big"
)

func addition(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subtract(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiplication(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func division(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}

func main() {
	a, _ := new(big.Int).SetString("777000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("999000000000000000000000000000000", 10)
	fmt.Printf("%v + %v = %v\n", a, b, addition(a, b))
	fmt.Printf("%v - %v = %v\n", a, b, subtract(a, b))
	fmt.Printf("%v * %v = %v\n", a, b, multiplication(a, b))
	fmt.Printf("%v / %v = %v\n", a, b, division(new(big.Float).SetInt(a), new(big.Float).SetInt(b)))
}
