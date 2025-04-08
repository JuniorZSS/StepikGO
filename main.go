package main

import (
	"fmt"
	"stepikgo/FOR"
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
		FOR.For1()

	case 2:
		FOR.For2()

	case 3:
		FOR.For3()

	case 4:
		FOR.For4()

	case 5:
		FOR.For5()

	case 6:
		FOR.For6()

	case 7:
		FOR.For7()

	case 100:
		FOR.Repeat()

	default:
		fmt.Println("Нет задачи с таким номером")
	}

}
