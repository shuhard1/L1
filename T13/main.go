package main

import "fmt"

//меняет местами a и b
func swap(a, b int) {
	fmt.Printf("a - %d b - %d\n", a, b)
	//a=23
	//b=45
	b = a + b //23+45=68
	a = b - a //68-23=45
	b = b - a //68-45=23
	//a=45
	//b=23
	fmt.Printf("a - %d b - %d\n", a, b)
}

func main() {
	swap(23, 45)
}
