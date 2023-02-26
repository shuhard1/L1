package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "snow dog sun"
	//закидываем каждое слово строки в слайс
	words := strings.Split(s, " ")
	//проходимся с последнего до первого элемента
	for i := range words {
		//вывод в stdout
		fmt.Print(words[len(words)-i-1], " ")
	}
}
