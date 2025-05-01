package __13FinalBasic

import "fmt"

func Task10() {
	var a, b int
	fmt.Print("Введите два числа, чтобы найти наибольшее число между ними кратное семи: ")
	fmt.Scan(&a, &b)
	c := b - ((b%7 + 7) % 7)
	if c > a || a == 0 {
		fmt.Println(c)
	} else {
		fmt.Println("NO")
	}
}
