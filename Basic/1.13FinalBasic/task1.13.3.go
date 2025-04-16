package __13FinalBasic

import "fmt"

func Task3() {
	var t, m, h int
	fmt.Print("Введите количество секунду и узнаете сколько это в минутах и часах: ")
	fmt.Scan(&t)
	if 0 < t && t <= 86399 {
		m = t / 60
		h = m / 60
		m = m % 60
		fmt.Printf("It is %v hours %v minutes.", h, m)
	} else {
		fmt.Println("Введены некорректные числа, введите числа не более 86399 и не меньше, и не равно нулю")
	}
}
