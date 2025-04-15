package __12ArraySlice

import (
	"fmt"
)

func Arr5() {
	fmt.Println("Задание сделать программу в которой на ввод дадут N количество элементов в массиве," +
		" \nдалее зададут значения для массива, а программа посчитает сколько положительных целых чисел в массиве")
	fmt.Print("Введите количество значений в массиве: ")

	var a int
	var j int
	fmt.Scan(&a)

	arr := make([]int, a)

	fmt.Print("Вводите массив целых чисел: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	for i := range arr {
		if arr[i] > 0 {
			j++
		}
	}
	fmt.Println(j)
}
