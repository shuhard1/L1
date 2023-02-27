package main

import (
	"log"
	"strconv"
)

// устанавливает i-й бит в 1 или 0
func changeBit(num, position, newBit int) int64 {
	//FormatInt возвращает строковое представление i в заданной системе счисления
	bits := []byte(strconv.FormatInt(int64(num), 2))
	newB := []byte(strconv.FormatInt(int64(newBit), 2))
	//заменяем элемент по индексу в массиве байтов на 1
	bits[position] = newB[0]
	//приводим массив byte в строку
	//получаем из строки int64
	newVal, err := strconv.ParseInt(string(bits), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return newVal
}
func main() {
	num := 2
	position := 1
	log.Println(changeBit(num, position, 1))
}
