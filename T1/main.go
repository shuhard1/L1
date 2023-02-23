package main

import "fmt"

//объявление структуры Human
type Human struct {
	message string
}

//объявление структуры Action
type Action struct {
	//поле является структурой типа Human
	Human
}

//создаем метод связанный с типом Human
func (h *Human) someAction() {
	fmt.Println(h.message)
}

func main() {
	//создаем экземпляр типа Action
	var action Action

	//обращаемя к полю структуры Human внутри структуры Action
	//присваиваем значение полю message
	action.message = "some action"

	//Все методы типа Human автоматически делаются доступными через тип Action
	action.someAction()
}
