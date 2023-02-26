package main

import (
	"fmt"
)

func showType(data interface{}) {
	fmt.Printf("type  %T\n", data)
}

func main() {
	showType("str")
	showType(33)
	showType(33.33)
	showType(make(chan int))
}
