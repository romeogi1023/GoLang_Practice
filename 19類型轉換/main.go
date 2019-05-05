package main

import "fmt"

func main() {

	a, b, c := 54, 30, 20

	sum := a + b + c
	//類型轉換 用來在不同但相互兼容的類型之間相互轉換的方式
	//1.數據類型（變量）  int（a）
	//2.數據類型（表達式）
	fmt.Println(float64(sum) / 3)
	//不兼容
	//var flag bool
	//flag = true
	//fmt.Println(int(flag))
	//
	//flag = bool(1)

	var d float32 = 3.1
	var f float64 = 3.5

	//在類型轉換時建議低類型轉成高類型  保證數據的精度
	//int8---int16---int32---int64
	//float32---float64
	//int64 ----float64()
	//建議整型轉成浮點型
	num := float64(d) + f
	fmt.Println(num)

	//高類型----低類型  會丟失精度  數據溢出  符號發生改變

	var g int = 1234
	fmt.Println(int8(g))

}
