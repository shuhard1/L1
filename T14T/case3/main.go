package main

import (
	"fmt"
	"reflect"
)

func showType(data interface{}) {
	fmt.Println("type ", reflect.ValueOf(data).Kind())
}

func main() {
	showType("str")
	showType(33)
	showType(33.33)
	showType(make(chan int))
}
