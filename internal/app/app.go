package app

import (
	"fmt"
	FOR "stepikgo/internal/mywork/stepic-tasks/Basic/1.10for"
	__11fmt "stepikgo/internal/mywork/stepic-tasks/Basic/1.11fmt"
	__12ArraySlice2 "stepikgo/internal/mywork/stepic-tasks/Basic/1.12ArraySlice"
	__13FinalBasic2 "stepikgo/internal/mywork/stepic-tasks/Basic/1.13FinalBasic"
	testTask1_13_10 "stepikgo/internal/mywork/stepic-tasks/Basic/1.13FinalBasic/testTask1.13.10"
	__1Func "stepikgo/internal/mywork/stepic-tasks/Basic/2.1Func"
)

type Task interface {
	Run()
}

type TaskFunc func()

func (f TaskFunc) Run() { f() }

var taskList = map[int]Task{
	1:   TaskFunc(FOR.For1),
	2:   TaskFunc(FOR.For2),
	3:   TaskFunc(FOR.For3),
	4:   TaskFunc(FOR.For4),
	5:   TaskFunc(FOR.For5),
	6:   TaskFunc(FOR.For6),
	7:   TaskFunc(FOR.For7),
	8:   TaskFunc(FOR.For8),
	9:   TaskFunc(__12ArraySlice2.Arr1),
	10:  TaskFunc(__12ArraySlice2.Arr2),
	11:  TaskFunc(__12ArraySlice2.Arr3),
	12:  TaskFunc(__12ArraySlice2.Arr4),
	13:  TaskFunc(__12ArraySlice2.Arr5),
	14:  TaskFunc(__13FinalBasic2.Task1),
	15:  TaskFunc(__13FinalBasic2.Task2),
	16:  TaskFunc(__13FinalBasic2.Task3),
	17:  TaskFunc(__13FinalBasic2.Task4),
	18:  TaskFunc(__13FinalBasic2.Task5),
	19:  TaskFunc(__13FinalBasic2.Task6),
	20:  TaskFunc(__13FinalBasic2.Task7),
	21:  TaskFunc(__13FinalBasic2.Task8),
	22:  TaskFunc(__13FinalBasic2.Task9),
	23:  TaskFunc(testTask1_13_10.Task10),
	24:  TaskFunc(__13FinalBasic2.Task11),
	25:  TaskFunc(__13FinalBasic2.Task12),
	26:  TaskFunc(__13FinalBasic2.Task13),
	27:  TaskFunc(__13FinalBasic2.Task14),
	28:  TaskFunc(__13FinalBasic2.Task15),
	29:  TaskFunc(__1Func.Task2),
	30:  TaskFunc(__1Func.Task3),
	31:  TaskFunc(__1Func.Task4),
	32:  TaskFunc(__1Func.Task5),
	101: TaskFunc(__11fmt.Format),
}

func Menu() {
	fmt.Println("Добро пожаловать в программу решения задач по STEPIK")
	fmt.Println("Выберите задачу: 1 - 100")
	fmt.Print("Введите номер задачи: ")
	var num int
	fmt.Scan(&num)
	if task, ok := taskList[num]; ok {
		task.Run()
	} else {
		fmt.Println("Нет задачи с таким номером")
	}
}
