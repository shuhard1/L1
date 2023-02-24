package main

import (
	"fmt"
	"time"
)

// Через закрытия канала
func main() {
	ch := make(chan string, 6)
	go case1(ch)
	go case2(ch)
	ch <- "hello"
	ch <- "world"
	//закрываем канал
	close(ch)
	//ждем, чтобы горутины успели остановиться
	time.Sleep(time.Second)
}

func case1(ch <-chan string) {
	for {
		v, ok := <-ch
		//проверка закрыт ли канал
		if !ok {
			fmt.Println("Finish")
			//выходим из функции и горутина останавливается
			return
		}
		fmt.Println(v)
	}
}

func case2(ch <-chan string) {
	//range будет считывать данные из канала до тех пор, пока канал не будет закрыт
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Finish")
}
