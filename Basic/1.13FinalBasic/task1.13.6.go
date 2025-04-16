package __13FinalBasic

import "fmt"

func Task6() {
	fmt.Print("Введите два числа и получите среднее арифметическое число: ")
	var a, b float64
	fmt.Scan(&a, &b)
	a = (a + b) / 2
	fmt.Print(a)
}
