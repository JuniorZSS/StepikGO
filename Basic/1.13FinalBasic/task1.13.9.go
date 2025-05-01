package __13FinalBasic

import "fmt"

func Task9() {
	fmt.Print("Введите число и получите его цифровой корень: ")
	var a, b, c int
	fmt.Scan(&a)
	for a > 0 || b > 0 {
		b += a % 10
		a /= 10
		//fmt.Println("Вывожу b из цикла ", b)
		if a == 0 {
			c += b % 10
			b /= 10
			//fmt.Println("Вывожу с из цикла ", c)
		}
	}

	fmt.Println(c)
}
