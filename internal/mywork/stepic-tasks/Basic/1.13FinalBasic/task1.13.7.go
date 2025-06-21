package __13FinalBasic

import "fmt"

func Task7() {
	fmt.Println("Создайте массив значений и узнайте сколько значений из него равны нулю")
	fmt.Print("Введите количество значений в массиве: ")

	var a int
	var j int
	fmt.Scan(&a)

	arr := make([]int, a)

	fmt.Print("Введите массив целые числа: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	for i := range arr {
		if arr[i] == 0 {
			j++
		}
	}
	fmt.Println(j)
}
