package main

import (
	"fmt"
	"stepikgo/Basic/ArrSlice"
	"stepikgo/Basic/FinalBasic"
	"stepikgo/Basic/Fmti"
	FOR2 "stepikgo/Basic/For"
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

	case 9:
		ArrSlice.Arr1()

	case 10:
		ArrSlice.Arr2()

	case 11:
		ArrSlice.Arr3()

	case 12:
		ArrSlice.Arr4()

	case 13:
		ArrSlice.Arr5()

	case 14:
		FinalBasic.Task1()

	case 15:
		FinalBasic.Task2()

	case 100:
		FOR2.Repeat()

	case 101:
		Fmti.Format()

	case 8:
		Fmti.Fmti1()

	default:
		fmt.Println("Нет задачи с таким номером")
	}

}
