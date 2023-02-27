package main

import (
	"fmt"
	"reflect"
)

// определяет тип с помощью TypeOf
func showType(data interface{}) {
	//TypeOf возвращает тип
	fmt.Println("type ", reflect.TypeOf(data))
}

func main() {
	showType("str")
	showType(123)
	showType(123.123)
	showType(make(chan int))
}
