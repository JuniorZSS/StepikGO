package main

import (
	"fmt"
	"stepikgo/src/Basic/1.10for"
	"stepikgo/src/Basic/1.11fmt"
	__12ArraySlice2 "stepikgo/src/Basic/1.12ArraySlice"
	__13FinalBasic2 "stepikgo/src/Basic/1.13FinalBasic"
	"stepikgo/src/Basic/1.13FinalBasic/testTask1.13.10"
)

func main() {

	fmt.Println("Добро пожаловать!" +
		"\nВсе задачи: 1-100")
	// Доработать текст!
	fmt.Print("Введите номер задачи которую нужно воспроизвести: ")

	var taskNum int
	fmt.Scan(&taskNum)

	// Реализовать интерфейс и роутинг задач, более эффективно, удобнее, менее затратно по памяти.
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

	case 8:
		FOR.For8()

	case 9:
		__12ArraySlice2.Arr1()

	case 10:
		__12ArraySlice2.Arr2()

	case 11:
		__12ArraySlice2.Arr3()

	case 12:
		__12ArraySlice2.Arr4()

	case 13:
		__12ArraySlice2.Arr5()

	case 14:
		__13FinalBasic2.Task1()

	case 15:
		__13FinalBasic2.Task2()

	case 16:
		__13FinalBasic2.Task3()

	case 17:
		__13FinalBasic2.Task4()

	case 18:
		__13FinalBasic2.Task5()

	case 19:
		__13FinalBasic2.Task6()

	case 20:
		__13FinalBasic2.Task7()

	case 21:
		__13FinalBasic2.Task8()

	case 22:
		__13FinalBasic2.Task9()

	case 23:
		testTask1_13_10.Task10()

	case 24:
		__13FinalBasic2.Task11()

	case 25:
		__13FinalBasic2.Task12()

	case 26:
		__13FinalBasic2.Task13()

	case 27:
		__13FinalBasic2.Task14()

	case 28:
		__13FinalBasic2.Task15()

	case 100:
		FOR.Repeat()

	case 101:
		__11fmt.Format()

	default:
		fmt.Println("Нет задачи с таким номером")
	}

}
