package __13FinalBasic

import "fmt"

func Task8() {
	fmt.Println("Создайте массив и найдите количество минимальных значений!")
	var a int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&a)

	arr := make([]int, a)
	fmt.Print("Введите количество числа в массив: ")

	m := arr[0]
	j := 0
	for i := range arr {
		fmt.Scan(&arr[i])
		if m > arr[i] {
			j = 1
		}
		if m == arr[i] {
			j++
		}
		//Это дичь... Я получаю лишнюю единицу
		if j == 0 {
			j = 1
		}
	}

	fmt.Print(j)
}
