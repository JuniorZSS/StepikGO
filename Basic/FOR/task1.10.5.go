package FOR

import "fmt"

func For5() {

	var n, c, d int
	fmt.Println("Задание тут: https://stepik.org/lesson/228263/step/7?unit=200796 \n" +
		"Введите, например 20 3 5, программа выдаст число кратное второму числу, но не кратное третьему")
	fmt.Print("Вводите ваши числа: ")
	fmt.Scan(&n, &c, &d)

	for i := 1; i <= n; i++ {
		if i%c != 0 {
			continue
		}
		if i%d == 0 {
			continue
		}
		fmt.Println(i)
		break
	}

}
