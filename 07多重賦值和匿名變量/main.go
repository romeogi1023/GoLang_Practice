package main

import (
	"fmt"
)

func test07() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	// a:=10
	// b:=20
	// c:=30
	a, b := 10, 20
	fmt.Println(a, b)

	// a = b
	// b = a
	// fmt.Println(a, b)

	var c int
	c = a
	a = b
	b = c
	fmt.Println(a, b)

	// i := 10
	// j := 20
	i, j := 10, 20
	i, j = j, i
	fmt.Println(i, j)

	//_匿名變量   丟棄數據不處理
	tmp, _ := 7, 8
	fmt.Println(tmp)

	// var d, e, f int
	_, e, f := test07()
	fmt.Println(e, f)
}
