package FMTI

import "fmt"

func Fmti1() {
	var a float64
	fmt.Scan(&a)
	if a <= 0 {
		a *= a
		fmt.Println("%.4f", a)
	}
	//} else fmt.Println("число -12.21 не подходит")
	fmt.Println("число -12.21 не подходит")

}
