package __13FinalBasic

import "fmt"

func Task15() {

	fmt.Print("Введите число и цифру, которую нужно убрать из вашего числа. Ваше число и цифра: ")

	var a, b int
	fmt.Scan(&a, &b)

	res := 0
	arr := []int{}
	for a > 0 {
		res = a % 10
		if res != b {
			arr = append(arr, res)
		}
		a = a / 10
	}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}
}
