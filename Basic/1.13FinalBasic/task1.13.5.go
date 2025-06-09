package __13FinalBasic

import "fmt"

const (
	existsMsg    = "Треугольник существует"
	notExistsMsg = "Треугольник не существует"
)

func Task5() {
	var a, b, c int
	fmt.Print("Введите длины сторон треугольника a, b, c, и узнаете существует ли такой треугольник: ")
	fmt.Scan(&a, &b, &c)
	
	result := existsMsg
	if !(a < b+c && b < a+c && c < a+b) {
		result = notExistsMsg
	}
	fmt.Println(result)
}
