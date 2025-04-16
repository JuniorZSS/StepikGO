package __13FinalBasic

import "fmt"

func Task4() {
	var a, b, c int
	fmt.Print("Введите длины сторон треугольника a, b, c,  и узнаете является ли он прямоугольным: ")
	fmt.Scan(&a, &b, &c)

	if c*c == (a*a)+(b*b) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
