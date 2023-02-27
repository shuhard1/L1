package main

import (
	"fmt"
	"math/big"
)

// сложение двух типов *big.Int
func addition(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// вычетание двух типов *big.Int
func subtract(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// умножение двух типов *big.Int
func multiplication(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// деление двух типов *big.Float
func division(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}

func main() {
	//инициализируем *big.Int
	a, _ := new(big.Int).SetString("777000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("999000000000000000000000000000000", 10)
	fmt.Printf("%v + %v = %v\n", a, b, addition(a, b))
	fmt.Printf("%v - %v = %v\n", a, b, subtract(a, b))
	fmt.Printf("%v * %v = %v\n", a, b, multiplication(a, b))
	//преобразовываем *big.Int в *big.Float
	fmt.Printf("%v / %v = %v\n", a, b, division(new(big.Float).SetInt(a), new(big.Float).SetInt(b)))
}
