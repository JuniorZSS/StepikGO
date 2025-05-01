package __13FinalBasic

import "fmt"

func Task13() {
	fmt.Println("Введите число, а функция скажет какое это число по счету в фибоначи, либо выдаст -1")
	fmt.Print("Ваше число: ")

	var a int
	fmt.Scan(&a)
	if a <= 1 {
		fmt.Println("-1")
		return
	}

	e := 40
	sum1, sum2 := 0, 1
	arr := make([]int, e)

	for i := 0; i < e; i++ {
		arr[i] = sum1
		sum1, sum2 = sum2, sum1+sum2
	}

	for i, v := range arr {
		if v == a {
			fmt.Println(i)
			break
		} else if v > a {
			fmt.Print("-1")
			break
		}
	}
	//fmt.Print(arr)
}
