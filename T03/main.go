package main

import (
	"fmt"
	"sync"
)

func sumArray(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	//блокирует выполнение функции, пока счетчик больше 0
	var wg sync.WaitGroup
	//увеличивает счетчик WaitGroup на заданное целочисленное значение
	wg.Add(len(nums))

	for i := 0; i < len(nums); i++ {
		go func(j int) {
			//перед закрытием функции уменьшаем счетчик WaitGroup на 1
			defer wg.Done()
			//возводим в квадрат элементы массива
			nums[j] *= nums[j]
		}(i)
		//когда созданные горутины получают управление, значение переменной i уже изменилось
		//поэтому надо передать копию i в функцию
	}
	//ждет пока счетчик будет равен 0
	wg.Wait()
	//выводит сумму всех элементов массива
	fmt.Println(sumArray(nums))
}
