package __11fmt

import "fmt"

func Fmti1() {
	var a float64
	fmt.Println("Задание: https://stepik.org/lesson/351787/step/3?unit=335741")
	fmt.Print("Введите число, больше 0, с плавающей точкой, получите его квадрат: ")
	fmt.Scan(&a)

	switch {
	case a <= 0:
		fmt.Printf("число %4.2f не подходит\n", a)

	case a > 10000:
		fmt.Printf("%e", a)
	default:
		fmt.Printf("%.4f", a*a)
	}
}
