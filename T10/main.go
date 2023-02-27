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

	// проходимся по слайсу numbers
	for _, number := range numbers {
		//проверяем каждый элемент слайса numbers
		//и добавляем в нужный слайс
		if number > 30.0 && number < 40.0 {
			PlusThirty = append(PlusThirty, number)
		} else if number > 20.0 && number < 30.0 {
			PlusTwenty = append(PlusTwenty, number)
		} else if number > 10.0 && number < 20.0 {
			PlusTen = append(PlusTen, number)
		} else if number < -20.0 && number > -30.0 {
			MinusTwenty = append(MinusTwenty, number)
		}
	}

	fmt.Printf("-20:%v\n", MinusTwenty)
	fmt.Printf("10:%v\n", PlusTen)
	fmt.Printf("20:%v\n", PlusTwenty)
	fmt.Printf("30:%v\n", PlusThirty)
}
