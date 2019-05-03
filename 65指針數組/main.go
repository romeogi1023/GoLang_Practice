package main

import (
	"fmt"
)

func main6501() {
	a := 10
	b := 20
	c := 30

	//指針數組
	var arr [3]*int = [3]*int{&a, &b, &c}

	fmt.Println(arr)

	*arr[1] = 200
	fmt.Println(b)

	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
	}
}

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	c := [3]int{7, 8, 9}

	// p := &a
	// fmt.Println(p)

	var arr [3]*[3]int = [3]*[3]int{&a, &b, &c}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	(*arr[1])[1] = 555
	fmt.Println(b)
}
