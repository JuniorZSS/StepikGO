package FuncStrPoint

import (
	"fmt"
)

func vote(x int, y int, z int) int {
	//print your code here

	if y == x {
		return x
	} else {
		return z
	}

}

func Task3() {
	fmt.Println("Функция принимающая 3 значения 0 или 1, и возвращает то значение которое повторяется более одного раза")
	fmt.Print("Введите 0 или 1, не более трех раз: ")

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	res := vote(a, b, c)
	fmt.Print("Число повторяется более одного раза: ", res)
}
