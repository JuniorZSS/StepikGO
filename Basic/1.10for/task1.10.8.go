package FOR

import "fmt"

func For8() {

	var a, b int
	fmt.Scan(&a, &b)

	ab := 0
	ba := 0
	bb := b
	bc := 1
	res := 0

	//arr := []int{}
	for a > 0 {
		ab = a % 10
		a = a / 10
		b = bb
		for b > 0 {
			ba = b % 10
			b = b / 10
			if ab == ba {
				res = ab*bc + res
				bc *= 10
			}
		}
	}

	for res > 0 {
		bc /= 10
		a = res / bc
		res = res % bc

		fmt.Printf("%d ", a)
	}
}

//for bc > 0 {
//	n := res / bc
//	fmt.Printf("%d ", n)
//	//res = res / 10
//	bc = bc / 10
//}
//for i := len(arr) - 1; i >= 0; i-- {
//	fmt.Print(arr[i])
//}
