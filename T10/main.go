package main

import (
	"fmt"
)

// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5
func main() {
	//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5
	numbers := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	PlusThirty := []float64{}
	PlusTen := []float64{}
	PlusTwenty := []float64{}
	MinusTwenty := []float64{}

	for _, number := range numbers { //цикл проходится по срезу numbers
		if number > 30.0 && number < 40.0 { //проверка элемента
			PlusThirty = append(PlusThirty, number)
		} else if number > 20.0 && number < 30.0 {
			PlusTwenty = append(PlusTwenty, number)
		} else if number > 10.0 && number < 20.0 {
			PlusTen = append(PlusTen, number)
		} else if number < -20.0 && number > -30.0 {
			MinusTwenty = append(MinusTwenty, number)
		}
	}
	fmt.Print("(-20:{")
	printSlice(MinusTwenty)

	fmt.Print("(10:{")
	printSlice(PlusTen)

	fmt.Print("(20:{")
	printSlice(PlusTwenty)

	fmt.Print("(30:{")
	printSlice(PlusThirty)
}

func printSlice(s []float64) {
	fmt.Printf("%v\n", s) // вывод среза в консоль
}
