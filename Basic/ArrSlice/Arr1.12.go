package ArrSlice

import "fmt"

func Arr1() {
	var workArray [10]uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	var a, b uint8
	
	for i := 0; i < 10; i++ {
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	fmt.Println(workArray)
}
