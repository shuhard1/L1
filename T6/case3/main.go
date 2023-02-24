package main

import (
	"context"
	"fmt"
	"time"
)

// Через context.WithCancel
func main() {
	ch := make(chan struct{})
	//Контекст позволяет распространять отмену по всей программе
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			//Done возвращает канал, который закрывается, когда Context отменяется
			//действует как сигнал отмены для функций
			case <-ctx.Done():
				//записываем пустую структуру в канал ch
				ch <- struct{}{}
				//останавливаем горутину
				return
			default:
				fmt.Println("foo...")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		//cancel() передается всем go-процедурам, которые вызывают .Done()
		cancel()
	}()

	//горутина main блокируется, пока не считает данные с канала ch
	<-ch
	fmt.Println("Finish")
}
