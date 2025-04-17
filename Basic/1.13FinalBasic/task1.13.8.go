package __13FinalBasic

import "fmt"

func Task8() {
	fmt.Println("Создайте массив и найдите количество минимальных значений!")
	var a int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&a)

	arr := make([]int, a)
	fmt.Print("Введите количество числа в массив: ")

	j := 1
	for i := range arr {
		m := arr[0]
		fmt.Scan(&arr[i])
		if m > arr[i] {
			j = 0
		}
		if m == arr[i] {
			j++
		}
	}

	fmt.Print(j)
}
