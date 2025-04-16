package __13FinalBasic

import "fmt"

func Task5() {
	var a, b, c int
	fmt.Print("Введите длины сторон треугольника a, b, c,  и узнаете существует ли такой треугольник: ")
	fmt.Scan(&a, &b, &c)
	if a < b+c && b < a+c && b < a+c {
		fmt.Print("Существует")
	} else {
		fmt.Print("Не существует")
	}
}
