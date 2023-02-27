package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	select {
	//After блокирует горутину, пока истечет продолжительность,
	//а затем отправляет текущее время по возвращенному каналу
	case <-time.After(duration):
	}
}

func main() {
	fmt.Println("start")
	sleep(5 * time.Second)
	fmt.Println("finish")
}
