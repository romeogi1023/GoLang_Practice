package main

import (
	"fmt"
)

func test86(i int) {
	var arr [3]int
	arr[i] = 999
	//panic("runtime error: index out of range")
	fmt.Println(arr)
}

func main() {
	// fmt.Println("hello")
	// fmt.Println("hello")
	// fmt.Println("hello")
	// //當程序遇到panic時，會自動終止
	// panic("hello")
	// fmt.Println("hello")
	// fmt.Println("hello")

	test86(3)
}
