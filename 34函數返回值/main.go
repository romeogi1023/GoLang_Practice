package main

import "fmt"

func sum(a, b int) (c, d, sum int) {
	c = 3
	d = 8

	sum = a + b
	return
}

func main() {
	var result int
	var result2 int
	var result3 int
	result2, result3, result = sum(5, 7)
	fmt.Println(result2, result3, result)

	var test int
	var test3 int
	//匿名變量
	_, test3, test = sum(2, 7)
	fmt.Println( test3, test)
}
