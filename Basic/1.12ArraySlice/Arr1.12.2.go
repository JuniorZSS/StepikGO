package __12ArraySlice

import "fmt"

func Arr2() {
	fmt.Printf("Введите \"5 41 -231 24 49 6\", выдаст 4й элемент среза: ")
	s := make([]int, 5)

	for i := range s {
		fmt.Scan(&s[i])
	}

	for _, r := range s[4:] {
		fmt.Printf("%v", r)
	}
}
