package main

import "fmt"

//指針作為函數參數
func Swap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
	fmt.Println(*num1, *num2)
}

func main() {
	a := 10
	b := 20
	//值傳遞

	//&變量  取地址操作 引用運算符
	//*指針變量  取值操作  解引用運算符
	Swap(&a, &b)
	fmt.Println(a, b)

	//地址傳遞
}
