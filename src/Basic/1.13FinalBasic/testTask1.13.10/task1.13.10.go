package testTask1_13_10

import "fmt"

func Task10() {
	var a, b int
	fmt.Print("Введите два числа, чтобы найти наибольшее число между ними кратное семи: ")
	fmt.Scan(&a, &b)
	res := task10(a, b)
	if res == -1 {
		fmt.Print("NO")
	} else {
		fmt.Print(res)
	}
}

func task10(a int, b int) int {
	c := b - ((b%7 + 7) % 7)
	if c > a || a == 0 {
		return c
	} else {
		return -1
	}
}
