package main

import (
	"fmt"
	"reflect"
)

func showType(data interface{}) {
	fmt.Println("type ", reflect.TypeOf(data))
}

func main() {
	showType("str")
	showType(33)
	showType(33.33)
	showType(make(chan int))
}
