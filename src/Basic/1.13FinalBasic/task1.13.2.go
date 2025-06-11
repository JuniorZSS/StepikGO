package __13FinalBasic

import "fmt"

func Task2() {
	fmt.Print("Введите трехзначное число а программа напишет его в обратном порядке: ")
	var a, b, k int
	fmt.Scan(&a)
	b = a / 100
	k = a % 10
	a = (a / 10) % 10

	arr := []int{k, a, b}

	for _, v := range arr {
		fmt.Print(v)
	}

}
