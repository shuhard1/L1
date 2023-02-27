package main

import (
	"fmt"
)

// определяет тип с помощью fmt.Printf()
func showType(data interface{}) {
	fmt.Printf("type  %T\n", data)
}

func main() {
	showType("str")
	showType(123)
	showType(123.123)
	showType(make(chan int))
}
