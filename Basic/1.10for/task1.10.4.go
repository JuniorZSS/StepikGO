package FOR

import "fmt"

func For4() {

	var a int
	var i = 0
	var x = 0

	fmt.Println("Задание: на вход данна  последовательность натуральных чисел, которая заканчивается нулём \nНужно подсчитать сколько раз встречается наибольшее число из последовательности")
	fmt.Print("Введите последовательность натуральных чисел \"(кроме 0) в конце добавьте (0)\": ")

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > x {
			x = a
			i = 1
		} else if a == x {
			i++
		}

	}
	fmt.Println(i)
}
