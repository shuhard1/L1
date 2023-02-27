package main

import (
	"fmt"
	"reflect"
)

// определяет тип с помощью reflect.ValueOf().Kind()
func showType(data interface{}) {
	fmt.Println("type ", reflect.ValueOf(data).Kind())
}

func main() {
	showType("str")
	showType(123)
	showType(123.123)
	showType(make(chan int))
}
