package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// постоянное чтение данных с канала
func read(jobs <-chan int, wg *sync.WaitGroup) {
	//перед закрытием функции уменьшаем счетчик WaitGroup на 1
	defer wg.Done()
	//канал jobs зацикливается до тех пор, пока он не будет закрыт
	for j := range jobs {
		fmt.Println("read ", j)
	}
	log.Println("Shutdown read")
}

// постоянная запись данных в канал
func sender(jobs chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	//перед закрытием функции уменьшаем счетчик WaitGroup на 1
	defer wg.Done()

	for i := 0; ; i++ {

		select {
		//Done возвращает канал, который закрывается, когда Context отменяется
		//действует как сигнал отмены для функций
		case <-ctx.Done():
			log.Println("Shutdown sender")
			close(jobs)
			//выход из функции
			return
		default:
			//отправляем значения в канал
			jobs <- i
			fmt.Println("send ", i)
			//переход на следующую итерацию
			continue
		}
	}
}

func main() {
	const N = 1
	//Контекст позволяет распространять отмену по всей программе
	ctx, cancel := context.WithCancel(context.Background())
	//в этот канал будут отправлять значения, а потом читать
	jobs := make(chan int)

	var wg sync.WaitGroup
	//увеличивает счетчик WaitGroup на заданное количество
	wg.Add(2)
	//запускаем горутины
	go sender(jobs, ctx, &wg)
	go read(jobs, &wg)

	//канал timer будет уведомлен через N секунд
	timer := time.NewTimer(N * time.Second)
	//Блокирует <-timer.C канал таймера C до тех пор, пока не отправит значение, указывающее, что таймер сработал
	<-timer.C
	//cancel() передается всем go-процедурам, которые вызывают .Done()
	cancel()

	//ждет пока счетчик будет равен 0
	wg.Wait()
}
