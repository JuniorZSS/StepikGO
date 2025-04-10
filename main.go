package main

import (
	"fmt"
	"stepikgo/Basic/FMTI"
	FOR2 "stepikgo/Basic/FOR"
)

func main() {

	fmt.Println("Добро пожаловать!" +
		"\nВсе задачи:")

	fmt.Println("№1 Циклы 1.10")
	fmt.Println("№2 Циклы 1.10")
	fmt.Println("№4 Циклы 1.10")
	fmt.Println("№5 Циклы 1.10")
	fmt.Println("№6 Циклы 1.10")
	fmt.Println("№7 Циклы 1.10")
	fmt.Println("№8 Пакет fmt 1.11")
	fmt.Print("Введите номер задачи которую нужно воспроизвести: ")

	var taskNum int
	fmt.Scan(&taskNum)

	switch taskNum {
	case 1:
		FOR2.For1()

	case 2:
		FOR2.For2()

	case 3:
		FOR2.For3()

	case 4:
		FOR2.For4()

	case 5:
		FOR2.For5()

	case 6:
		FOR2.For6()

	case 7:
		FOR2.For7()

	case 100:
		FOR2.Repeat()

	case 8:
		FMTI.Fmti1()

	default:
		fmt.Println("Нет задачи с таким номером")
	}

}
