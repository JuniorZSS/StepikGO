package __1Func

import (
	"fmt"
	"time"
)

func minimumFromFour() int {
	//print your code here
	fmt.Println("Вводится четыре числа. Возвращается наименьшее из них")
	fmt.Print("Введите 4 числа: ")

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	res := a
	if res > b {
		res = b
	}
	if res > c {
		res = c
	}
	if res > d {
		res = d
	}
	return res
}

func Task2() {
	res := minimumFromFour()
	fmt.Print("Наименьшее число: ", res)
	time.Sleep(1 * time.Second) // Тестил задержку
}
