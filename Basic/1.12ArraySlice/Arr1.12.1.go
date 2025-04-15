package __12ArraySlice

import "fmt"

func Arr1() {
	var workArray [10]uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	var a, b uint8

	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for _, elem := range workArray {
		fmt.Printf("%d ", elem)
	}
}
