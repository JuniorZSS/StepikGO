package __13FinalBasic

import "fmt"

func Task14() {
	fmt.Println("Функция переводит введенное число в двоичную систему исчисления")
	fmt.Print("Введите число: ")

	var a int
	fmt.Scan(&a)

	for i := 0; a <= 0; i++ {
		arr[i] = a % 2
	}

}
