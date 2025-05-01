package __13FinalBasic

import "fmt"

func Task12() {
	fmt.Print("Задача, введите число и получите все числа, квадрат которых будет 2, до введенного вами числа(например 50), ваше число:")

	var a int
	fmt.Scan(&a)

	arr := make([]int, a)

	for i := 1; i <= a; {
		if i < 2 {
			arr = append(arr, 1)
		}
		i = i * 2
		if i > a {
			break
		}
		arr = append(arr, i)
	}

	for _, elem := range arr {
		if elem == 0 {
			continue
		} else {
			fmt.Printf("%d ", elem)
		}
	}
}
