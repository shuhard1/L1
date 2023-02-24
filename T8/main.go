package main

import (
	"fmt"
	"log"
	"strconv"
)

func IntToBits(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func bitsToInt(s string) int64 {
	newVal, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return newVal
}

func changeOneBit(num, position, newBit int) int64 {
	if position > 0 && num == 0 {
		return int64(num)
	}

	if newBit > 1 || newBit < 0 {
		log.Println("Новый бит может быть только 1 или 0")
		return int64(num)
	}

	bits := []byte(IntToBits(num))
	fmt.Printf("Было \t %s(%v)\n", string(bits), num)

	if position >= len(bits) {
		log.Printf("В числе %v нет %v разрядов. Разярядов в числе %v только %v\n", num, position, num, len(bits))
		return int64(num)
	}

	newB := []byte(IntToBits(newBit))
	bits[position] = newB[0]
	newVal := bitsToInt(string(bits))
	fmt.Printf("Стало \t %s(%v)\n", string(bits), newVal)
	return newVal
}
func main() {
	log.Println("value")
	var num int
	fmt.Scanln(&num)
	log.Println("index")
	var position int
	fmt.Scanln(&position)

	log.Println(changeOneBit(num, position, 1)) //8
}
