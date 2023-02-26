package main

import "fmt"

//Бинарный поиск — это стратегия поиска, используемая для поиска элементов в списке путем последовательного уменьшения объема искомых данных и, таким образом, увеличения скорости нахождения искомого термина.
//Чтобы использовать алгоритм бинарного поиска, список, над которым нужно работать, должен быть уже отсортирован.
func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
