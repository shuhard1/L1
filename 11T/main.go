package main

import (
	"fmt"
)

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{2, 3, 5, 7, 11}
	fmt.Println(Intersection(a, b))
}
