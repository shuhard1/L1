package main

import (
	"fmt"
	"strings"
)

// если использовать символы unicode,
// то string скопируется некорректно
// следует добавить преобразование в тип rune
var justString string

func createHugeString(count int) string {
	return strings.Repeat("👉👈", count)
}

func someFunc() {
	v := createHugeString(1 << 10)
	//преобразование в тип rune
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
