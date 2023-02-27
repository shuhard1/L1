package main

import "fmt"

//Адаптер действует как оболочка между двумя объектами.
//Он перехватывает обращения к одному объекту
//и преобразовывает их в формат и интерфейс, распознаваемый вторым объектом.
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
