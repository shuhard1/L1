package main

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstCh := make(chan int)
	secondCh := make(chan int)

	//в firstCh пишутся числа из массива nums
	go func() {
		for _, x := range nums {
			firstCh <- x
		}
		close(firstCh)
	}()

	go func() {
		//читает из канала, пока он не будет закрыт
		for x := range firstCh {
			//в secondCh пишется результат операции x*2
			secondCh <- x * 2
		}
		close(secondCh)
	}()

	//читает из канала, пока он не будет закрыт
	for n := range secondCh {
		//вывод в stdout
		println(n)
	}
}
