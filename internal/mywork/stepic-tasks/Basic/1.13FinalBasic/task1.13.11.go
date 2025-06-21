package __13FinalBasic

import "fmt"

func Task11() {
	var n int
	fmt.Println("Введи число, получите количество Коров :D")
	fmt.Print("И ваше число: ")
	fmt.Scan(&n)
	switch d := n % 100; {
	case d >= 11 && d <= 14:
		fmt.Printf("%d korov\n", n)
	default:
		switch n % 10 {
		case 1:
			fmt.Printf("%d korova\n", n)
		case 2, 3, 4:
			fmt.Printf("%d korovy\n", n)
		default:
			fmt.Printf("%d korov\n", n)
		}
	}
}
