package main

import (
	"fmt"
	FOR2 "stepikgo/Basic/1.10for"
	"stepikgo/Basic/1.11fmt"
	"stepikgo/Basic/1.12ArraySlice"
	"stepikgo/Basic/1.13FinalBasic"
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
		__12ArraySlice.Arr1()

	case 10:
		__12ArraySlice.Arr2()

	case 11:
		__12ArraySlice.Arr3()

	case 12:
		__12ArraySlice.Arr4()

	case 13:
		__12ArraySlice.Arr5()

	case 14:
		__13FinalBasic.Task1()

	case 15:
		__13FinalBasic.Task2()

	case 16:
		__13FinalBasic.Task3()

	case 17:
		__13FinalBasic.Task4()

	case 18:
		__13FinalBasic.Task5()

	case 19:
		__13FinalBasic.Task6()

	case 20:
		__13FinalBasic.Task7()

	case 21:
		__13FinalBasic.Task8()

	case 22:
		__13FinalBasic.Task9()

	case 23:
		__13FinalBasic.Task10()

	case 24:
		__13FinalBasic.Task11()

	case 25:
		__13FinalBasic.Task12()

	case 26:
		__13FinalBasic.Task13()

	case 27:
		__13FinalBasic.Task14()

	case 28:
		__13FinalBasic.Task15()

	case 100:
		FOR2.Repeat()

	case 101:
		__11fmt.Format()

	case 8:
		__11fmt.Fmti1()

	default:
		fmt.Println("Нет задачи с таким номером")
	}

}
