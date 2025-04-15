package __12ArraySlice

import "fmt"

func Arr3() {

	array := [5]int{}
	var a int
	fmt.Print("Введите массив из 5 целых чисел придет ответ с наибольшим из них: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	// ...
	for _, m := range array {
		if m > a {
			a = m
		}
	}
	fmt.Println(a)

}
