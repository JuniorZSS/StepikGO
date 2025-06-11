package __13FinalBasic

import "fmt"

func Task14() {
	fmt.Println("Функция переводит введенное число в двоичную систему исчисления")
	fmt.Print("Введите число: ")
	//
	//var a int
	//fmt.Scan(&a)
	//
	//if a == 0 {
	//	fmt.Println(a)
	//}
	//
	//arr := []int{}
	//
	//var b int
	//for a > 0 {
	//	b = a % 2
	//	a = a / 2
	//	arr = append(arr, b)
	//}
	//
	//for i := len(arr) - 1; i >= 0; i-- {
	//	fmt.Print(arr[i])
	//}

	//Оптимально
	var number int
	fmt.Scan(&number)
	fmt.Printf("%b", number)

	//Тоже вариант
	//for ; n > 0; n /= 2 {
	//	defer fmt.Print(n%2)
	//}

}
