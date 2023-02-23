package main

import (
	"fmt"
)

// main работает в своей собственной горутине
// после ее завершения все остальные горутины так же завершаются
func main() {
	nums := []int{2, 4, 6, 8, 10}

	//пробегаюсь по массиву
	for i := 0; i < len(nums); i++ {
		//создаю горутину ключевым словом go
		//выведет квадраты в stdout
		go fmt.Print(nums[i]*nums[i], " ")
	}

	//чтобы программа не завершилась
	var input string
	fmt.Scanln(&input)
}
