package main

import (
	"fmt"
)

// возвращает массив с общими элементами, присутствующие в a и b
func intersection(a, b []int) (c []int) {
	m := make(map[int]struct{})

	//ключ не может повторятся, поэтому
	//так можно получить только уникальные элементы в слайсе
	for _, item := range a {
		m[item] = struct{}{}
	}

	for _, item := range b {
		//если такой ключ есть в мапе, то
		//добавляем item в слайс c
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{2, 3, 5, 7, 11}
	fmt.Println(intersection(a, b))
}
