package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, shutdownChannel chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		//вывод в stdout
		fmt.Println("worker id", id, "started  job", j)
		//имитация работы
		time.Sleep(time.Second)
		select {
		case <-shutdownChannel:
			log.Println("Shutdown goroutine: id", id)
			return
		default:
			continue
		}
	}
}

func sender(jobs chan<- int) {
	for i := 0; ; i++ {
		jobs <- i
	}
}

func main() {
	var numWorkers int
	fmt.Scanln(&numWorkers)

	//создаем канал, который будет ловить сигнал с именем SIGINT
	quitChannel := make(chan os.Signal)
	//Notify заставляет пакет signal ретранслировать входящие сигналы в канал quitChannel
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	shutdownChannel := make(chan struct{})

	jobs := make(chan int, numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	go sender(jobs)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, shutdownChannel, &wg)
	}

	<-quitChannel // received SIGINT or SIGTERM
	close(shutdownChannel)

	wg.Wait()
}
