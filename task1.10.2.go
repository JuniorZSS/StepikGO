package main

import "fmt"

func For2() {
	fmt.Println("Задание: напишите программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.")
	fmt.Print("Введите два натуральных числа (В задаче на вход подается 1 и 5, можно через пробел): ")
	var a, b int
	fmt.Scan(&a, &b)
	sum := 0

	for i := a; i <= b; i++ {

		sum += i

	}
	fmt.Println(sum)
}
