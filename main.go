package main

import (
	"fmt"
)

func main() {

	fmt.Println("Добро пожаловать!" +
		"\nВсе задачи:")
	fmt.Println("Все задачи:")
	fmt.Println("№1 Циклы 1.10")
	fmt.Println("№2 Циклы 1.10")
	fmt.Println("№3 Циклы 1.10")
	fmt.Print("Введите номер задачи которую нужно воспроизвести: ")

	var taskNum int
	fmt.Scan(&taskNum)

	switch taskNum {
	case 1:
		For1()

	case 2:
		For2()

	case 3:
		For3()

	case 4:
		For4()

	default:
		fmt.Println("Нет задачи с таким номером")
	}

}
