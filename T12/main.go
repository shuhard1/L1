package main

import "fmt"

func newSet(slice []string) []string {
	resMap := make(map[string]bool)
	result := []string{}

	for _, key := range slice {
		resMap[key] = true
	}
	//ключ не может повторятся - в этом прикол
	//так можно найти повторы в слайсе
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
