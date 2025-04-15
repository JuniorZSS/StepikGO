package FinalBasic

import "fmt"

func Task1() {
	fmt.Print("Введите трехзначное число и получите сумму его цифр: ")
	var a, b, k int
	fmt.Scan(&a)
	b = a / 100
	k = a % 10
	a = (a / 10) % 10
	fmt.Println(a + b + k)
}
