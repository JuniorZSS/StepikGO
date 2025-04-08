package FOR

import "fmt"

func Repeat() {
	var arr []int
	var a int
	fmt.Println("Ага, попались?! Это моя секретная задачка, эксперементировал с мапами и срезами\n Вы вводите некоторое количество чисел и увидите сколько раз повторились те или иные числа, которые вы ввели\n Обязательно в конце укажите 0, иначе программа не выполнится :)")
	fmt.Print("Вводите числа: ")
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		arr = append(arr, a)
	}
	//var a = 5
	//fmt.Print("Введите данные в слайс: ")
	//fmt.Scan(a)
	//a := [5]int{0, 4, 10, 10, 10}
	//slice := make([]int, a)
	m := make(map[int]int)
	for _, n := range arr {
		m[n]++
	}
	for k, v := range m {
		fmt.Println("Число в массиве: ", k, ",", "повторилось, раз: ", v)
	}
	fmt.Println("Вот такой массив вы ввели: ", arr)

}
