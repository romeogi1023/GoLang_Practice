package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//uint無符號整型  int有符號整型
	//var a uint = -10
	//fmt.Println(a)

	b := 20
	fmt.Println(b)
	//%T打印變量所屬類型
	fmt.Printf("%T\n", b)
	fmt.Println(unsafe.Sizeof(b))

	var c int32 = 10

	fmt.Printf("%T\n", c)
	fmt.Println(unsafe.Sizeof(c))

}
