package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	var countWorkers int = 300

	//увеличивает счетчик WaitGroup на заданное количество
	wg.Add(countWorkers)
	for i := 0; i < countWorkers; i++ {
		//запускаем 300 горутин
		go func(i int) {
			//Блокировка, чтобы только одна горутина могла получить доступ к count
			mx.Lock()
			//инкремент
			count++
			mx.Unlock()
			wg.Done()
			//уменьшаем счетчик WaitGroup на 1
		}(i)
		//когда созданные горутины получают управление, значение переменной i уже изменилось
		//поэтому надо передать копию i в функцию
	}
	//ждет пока счетчик будет равен 0
	wg.Wait()

	fmt.Println(count)
}
