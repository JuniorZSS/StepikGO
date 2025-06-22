package __1Func

import (
	"fmt"
)

func fibonacci(n int) int {
	//print your code here

	sum1, sum2 := 0, 1
	arr := make([]int, n+1)

	for i := 0; i < n; i++ {
		arr[i] = sum1
		sum1, sum2 = sum2, sum1+sum2
	}
	return sum1
}

func Task4() {
	fmt.Println("Функция фибоначи, введите число от нуля и получите число из ряда фибоначи")
	fmt.Print("Ваше число: ")

	var res int
	fmt.Scan(&res)
	a := fibonacci(res)
	fmt.Println(a)
}
