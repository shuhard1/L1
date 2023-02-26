package main

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstCh := make(chan int)
	secondCh := make(chan int)

	go func() {
		for _, n := range nums {
			firstCh <- n
		}
		close(firstCh)
	}()

	go func() {
		for n := range firstCh {
			secondCh <- n * 2
		}
		close(secondCh)
	}()

	for n := range secondCh {
		println(n)
	}
}
