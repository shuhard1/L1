package main

import (
	"fmt"
	"time"
)

// Через канал с пустой структурой
func main() {
	ch := make(chan string, 6)
	//cоздаем канал структуры make(chan struct{})(пустая структура не требует памяти).
	done := make(chan struct{})

	go func() {
		for {
			select {
			//записывает в канал
			case ch <- "foo":
			//читает канал done
			case <-done:
				//закрываем канал ch
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//ждет 3 секунды, перед записью в канал done
	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	//читает канал ch, пока он открыт
	for i := range ch {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Finish")
}
