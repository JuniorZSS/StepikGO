package ArrSlice

import "fmt"

func Arr3() {

	array := [5]int{}
	var a int
	var i int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	// ...
	for _, m := range array {
		if m >= a {
			i = a
		}
	}
	fmt.Println(i)

}
