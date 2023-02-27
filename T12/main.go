package main

import "fmt"

func newSet(slice []string) []string {
	resMap := make(map[string]struct{})
	result := []string{}
	//ключ не может повторятся, поэтому
	//так можно получить только уникальные элементы в слайсе
	for _, key := range slice {
		resMap[key] = struct{}{}
	}
	//записываем ключи resMap в слайс
	for key := range resMap {
		result = append(result, key)
	}

	return result
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := newSet(slice)
	fmt.Println(set)
}
