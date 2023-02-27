package main

import "fmt"

//не сохраняет порядок
func main() {
	sl := []string{"A", "B", "C", "D", "E"}
	i := 2

	//копирую последний элемент в a[i]
	sl[i] = sl[len(sl)-1] //{"A", "B", "E", "D", "E"}
	//Обрезаем последний элемент в слайсе
	sl = sl[:len(sl)-1] //{"A", "B", "E", "D"}

	fmt.Println(sl)
}
