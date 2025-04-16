package __13FinalBasic

import "fmt"

func Task8() {
	fmt.Println("Создайте массив и найдите количество минимальных значений!")
	var a int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&a)

	arr := make([]int, a)
	fmt.Print("Введите количество числа в массив: ")

	for i := range arr {
		fmt.Scan(&arr[i])
	}
	m := arr[0]
	for i := range arr {
		if m > arr[i] {
			m = arr[i]
		}
	}
	j := 0
	for i := range arr {
		if m == arr[i] {
			j++
		}
	}
	fmt.Print(j)
}
