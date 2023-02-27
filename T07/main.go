package main

import (
	"fmt"
	"sync"
	"time"
)

func setMap(m map[int]int, mx *sync.Mutex, key int, value int) {
	//Блокировка, чтобы только одна горутина могла получить доступ к карте
	mx.Lock()
	//запись
	m[key] = value
	mx.Unlock()
}

func main() {
	//инициализация
	m := make(map[int]int)
	var mx sync.Mutex

	for i := 0; i < 100; i++ {
		//запускаем горутины, которые будут конкурентно заполнять мапу
		go setMap(m, &mx, i, i)
	}

	for i, k := range m {
		fmt.Println("key - ", k, "value  - ", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}
