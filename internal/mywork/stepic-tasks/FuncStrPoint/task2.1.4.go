package FuncStrPoint

import "fmt"

func sumInt(a ...int) (int, int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return len(a), sum
}

func Task5() {
	fmt.Print("ВАЖНО, на момент написания этой таски, она решена, но я устал и не смог реализовать чтение N введенных значений и передачу их в функцию.")
	fmt.Print("Функции задача с множеством принимаемых значений в функции")
	fmt.Print("Тест введи до 10 чисел: ")
	//arr := make([]int)
	//fmt.Scan(&arr)

	// Вводим любое значение ниже
	count, res := sumInt(4, 5, 2, 424, 53, 5, 2, 4)
	fmt.Println(count, res)

}
