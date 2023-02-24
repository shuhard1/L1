package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	//перед закрытием функции уменьшаем счетчик WaitGroup на 1
	defer wg.Done()
	for j := range jobs {
		//вывод в stdout
		fmt.Println("worker id", id, "started  job", j)
		//имитация работы
		time.Sleep(time.Second)
		select {
		//Done возвращает канал, который закрывается, когда Context отменяется
		//действует как сигнал отмены для функций
		case <-ctx.Done():
			log.Println("Shutdown goroutine: id", id)
			//выход из функции
			return
		default:
			//переход на следующую итерацию
			continue
		}
	}
}

// постоянная запись данных в канал
func sender(jobs chan<- int) {
	for i := 0; ; i++ {
		jobs <- i
	}
}

func main() {
	//возможность выбора количества воркеров при старте
	var numWorkers int
	fmt.Scanln(&numWorkers)

	//создаем канал, который будет ловить сигнал с именем SIGINT
	quitChannel := make(chan os.Signal)
	//Notify заставляет пакет signal ретранслировать входящие сигналы в канал quitChannel
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	//Контекст позволяет распространять отмену по всей программе
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan int, numWorkers)

	var wg sync.WaitGroup
	//увеличивает счетчик WaitGroup на количество воркеров
	wg.Add(numWorkers)

	go sender(jobs)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, ctx, &wg)
	}

	// получает SIGINT
	<-quitChannel
	//cancel() передается всем go-процедурам, которые вызывают .Done()
	cancel()

	//ждет пока счетчик будет равен 0
	wg.Wait()
}
