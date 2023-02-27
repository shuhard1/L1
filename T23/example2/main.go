package main

import "fmt"

//сохраняет порядок
func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	//copy копирует элементы из исходного слайса в целевой слайс
	//a[i:] - {"C", "D", "E"}
	//a[i+1:] - {"D", "E"}
	copy(a[i:], a[i+1:]) //{"A", "B", "D", "E", "E"}
	//Обрезаем последний элемент в слайсе
	a = a[:len(a)-1]

	fmt.Println(a)
}
