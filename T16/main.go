package main

import "fmt"

//проверить каждый элемент массива по отношению к следующему элементу
//чтобы увидеть, больше ли он, если да, то нужно поменять их местами.
func main() {
	var sl = []int{1, 39, 2, 9, 7, 54, 11}

	var isDone bool

	for !isDone {
		isDone = true
		i := 0
		for i < len(sl)-1 {
			//если sl[i] больше следующего элемента, то меняем местами
			//таким образом больший элемент двигается вперед по слайсу
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
				isDone = false
			}
			i++
		}
	}

	fmt.Println(sl)
}
