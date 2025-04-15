package __12ArraySlice

import "fmt"

func Arr4() {
	fmt.Println("Задача: \"https://stepik.org/lesson/228265/step/15?auth=login&unit=200798\"")

	fmt.Print("Задайте длину массива: ")
	var a int
	fmt.Scan(&a)

	var array []int = make([]int, a)
	var result []int = make([]int, 0)

	fmt.Print("Вводите массив: ")

	for i := range array {
		fmt.Scan(&array[i])
		if array[i] > 0 && array[i] <= 100 && i%2 == 0 {
			result = append(result, array[i])
		}
	}
	//Выводим значения массива без квадратных скобок чтоб его...
	for _, elem := range result {
		fmt.Printf("%d ", elem)
	}
}
