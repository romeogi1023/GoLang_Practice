package main

import "fmt"

func main() {

	var t complex64
	t = 2.1 + 3.14i
	fmt.Println(t)

	//自動推導類型
	//默認為complex128
	t1 := 3.3 + 4.4i
	fmt.Printf("%T\n", t1)

	//通過內建函數  取實部   虛部
	fmt.Println(real(t1), imag(t1))

}
