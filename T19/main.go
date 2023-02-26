package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "главрыба"
	//закидываем каждый символ строки в слайс
	sl := strings.Split(s, "")
	//проходимся с последнего до первого элемента
	for i := range sl {
		//вывод в stdout
		fmt.Print(sl[len(sl)-1-i])
	}
}
